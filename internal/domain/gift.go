package domain

import (
	"fmt"
	"strings"
)

// GiftItem представляет предмет в каталоге подарков.
type GiftItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Weight   float64 `json:"weight"`
	MinAge   int     `json:"min_age"`

	// Metadata содержит дополнительную информацию о предмете
	// Используется для проверки соответствия специальным требованиям
	Metadata GiftMetadata `json:"metadata,omitempty"`
}

// GiftMetadata содержит метаданные для проверки соответствия требованиям.
type GiftMetadata struct {
	// Dietary информация
	ContainsMeat    bool `json:"contains_meat,omitempty"`
	ContainsFish    bool `json:"contains_fish,omitempty"`
	ContainsDairy   bool `json:"contains_dairy,omitempty"`
	ContainsNuts    bool `json:"contains_nuts,omitempty"`
	ContainsGluten  bool `json:"contains_gluten,omitempty"`
	ContainsSugar   bool `json:"contains_sugar,omitempty"`
	SugarFree       bool `json:"sugar_free,omitempty"`
	Vegetarian      bool `json:"vegetarian,omitempty"`
	Vegan           bool `json:"vegan,omitempty"`
	HalalCertified  bool `json:"halal_certified,omitempty"`
	KosherCertified bool `json:"kosher_certified,omitempty"`

	// Safety информация
	HasSmallParts  bool    `json:"has_small_parts,omitempty"`
	SmallPartsSize float64 `json:"small_parts_size,omitempty"` // в см
	Hypoallergenic bool    `json:"hypoallergenic,omitempty"`
	NonToxic       bool    `json:"non_toxic,omitempty"`
	Washable       bool    `json:"washable,omitempty"`
	FlameRetardant bool    `json:"flame_retardant,omitempty"`
	BPAFree        bool    `json:"bpa_free,omitempty"`

	// Medical информация
	HasFlashingLights  bool `json:"has_flashing_lights,omitempty"`
	HasFuzzyMaterial   bool `json:"has_fuzzy_material,omitempty"`
	IsDusty            bool `json:"is_dusty,omitempty"`
	CalmingEffect      bool `json:"calming_effect,omitempty"`
	Tactile            bool `json:"tactile,omitempty"`
	Predictable        bool `json:"predictable,omitempty"`
	WirelessCompatible bool `json:"wireless_compatible,omitempty"`
	AccessibleSize     bool `json:"accessible_size,omitempty"`

	// Other информация
	EcoFriendly      bool `json:"eco_friendly,omitempty"`
	Educational      bool `json:"educational,omitempty"`
	GenderNeutral    bool `json:"gender_neutral,omitempty"`
	Bilingual        bool `json:"bilingual,omitempty"`
	Durable          bool `json:"durable,omitempty"`
	Repairable       bool `json:"repairable,omitempty"`
	CharitySupported bool `json:"charity_supported,omitempty"`

	// Дополнительные поля
	Materials      []string `json:"materials,omitempty"`
	Certifications []string `json:"certifications,omitempty"`
	Warnings       []string `json:"warnings,omitempty"`
}

// GiftCategory представляет категорию подарков.
type GiftCategory struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	BaseWeight  float64 `json:"base_weight,omitempty"`
	MinAge      int     `json:"min_age,omitempty"`
	MaxAge      int     `json:"max_age,omitempty"`
}

// GiftRepository определяет интерфейс для работы с каталогом подарков.
type GiftRepository interface {
	// FindByID возвращает подарок по идентификатору.
	FindByID(id int) (*GiftItem, error)

	// FindByCategory возвращает подарки по категории.
	FindByCategory(category string) ([]GiftItem, error)

	// FindAll возвращает все подарки.
	FindAll() ([]GiftItem, error)

	// FindCheaperAlternative ищет более дешевую альтернативу.
	FindCheaperAlternative(item GiftItem, maxPrice float64) (*GiftItem, error)

	// FindByAgeRange возвращает подарки для указанного возрастного диапазона.
	FindByAgeRange(minAge, maxAge int) ([]GiftItem, error)

	// FindByRequirements возвращает подарки, соответствующие требованиям.
	FindByRequirements(requirements *SpecialRequirements) ([]GiftItem, error)
}

// GiftCatalogJSON представляет структуру JSON файла каталога подарков.
type GiftCatalogJSON struct {
	Version     string         `json:"version,omitempty"`
	GeneratedAt string         `json:"generated_at,omitempty"`
	Description string         `json:"description,omitempty"`
	Categories  []GiftCategory `json:"categories"`
	Items       []GiftItem     `json:"items"`
}

// CompliesWithDietary проверяет соответствие подарка диетическому требованию.
// Возвращает true если подарок соответствует требованию.
func (g *GiftItem) CompliesWithDietary(req DietaryRequirement) bool {
	switch req {
	case DietaryVegetarian:
		// Вегетарианство: исключить продукты с мясом, рыбой
		return !g.Metadata.ContainsMeat && !g.Metadata.ContainsFish || g.Metadata.Vegetarian

	case DietaryVegan:
		// Веганство: исключить все продукты животного происхождения
		return (!g.Metadata.ContainsMeat && !g.Metadata.ContainsFish &&
			!g.Metadata.ContainsDairy) || g.Metadata.Vegan

	case DietaryNutsAllergy:
		// Аллергия на орехи: исключить орехи и следы орехов
		return !g.Metadata.ContainsNuts

	case DietaryLactoseIntolerant:
		// Непереносимость лактозы: исключить молочные продукты
		return !g.Metadata.ContainsDairy

	case DietaryGlutenFree:
		// Без глютена: исключить пшеницу, ячмень, рожь
		return !g.Metadata.ContainsGluten

	case DietaryDiabetes:
		// Диабет: ограничить сахар, использовать сахарозаменители
		return !g.Metadata.ContainsSugar || g.Metadata.SugarFree

	case DietaryHalal:
		// Халяль: соответствие исламским нормам
		return g.Metadata.HalalCertified ||
			(!g.ContainsPork() && g.IsSlaughteredAccordingToHalal())

	case DietaryKosher:
		// Кошерность: соответствие иудейским нормам
		return g.Metadata.KosherCertified || g.IsKosherByIngredients()

	default:
		return true // Неизвестное требование - пропускаем
	}
}

// CompliesWithSafety проверяет соответствие подарка требованию безопасности.
// Возвращает true если подарок соответствует требованию.
func (g *GiftItem) CompliesWithSafety(req SafetyRequirement) bool {
	switch req {
	case SafetyNoSmallParts:
		// Без мелких деталей: детали > 3см для детей до 3 лет
		return !g.Metadata.HasSmallParts || g.Metadata.SmallPartsSize >= 3.0

	case SafetyHypoallergenic:
		// Гипоаллергенность: без аллергенов в материалах
		return g.Metadata.Hypoallergenic ||
			!g.ContainsAllergenicMaterials()

	case SafetyNonToxic:
		// Нетоксичность: сертифицированные безопасные материалы
		return g.Metadata.NonToxic ||
			g.HasSafetyCertification()

	case SafetyWashable:
		// Моющийся: возможность стирки/мытья
		return g.Metadata.Washable

	case SafetyFlameRetardant:
		// Огнестойкость: трудновоспламеняемые материалы
		return g.Metadata.FlameRetardant

	case SafetyBPAFree:
		// Без BPA: без бисфенола-А в пластике
		return g.Metadata.BPAFree ||
			!g.ContainsBPA()

	default:
		return true // Неизвестное требование - пропускаем
	}
}

// CompliesWithMedical проверяет соответствие подарка медицинскому требованию.
// Возвращает true если подарок соответствует требованию.
func (g *GiftItem) CompliesWithMedical(req MedicalRequirement) bool {
	switch req {
	case MedicalEpilepsy:
		// Эпилепсия: избегать мигающих светодиодов
		return !g.Metadata.HasFlashingLights

	case MedicalAsthma:
		// Астма: избегать пушистых/пылящих материалов
		return !g.Metadata.HasFuzzyMaterial && !g.Metadata.IsDusty

	case MedicalADHDFriendly:
		// Для СДВГ: успокаивающие, фокусирующие игрушки
		return g.Metadata.CalmingEffect ||
			g.IsFocusEnhancing()

	case MedicalAutismFriendly:
		// Для аутизма: тактильные, предсказуемые игрушки
		return (g.Metadata.Tactile || g.Metadata.Predictable) ||
			g.IsAutismFriendlyByDesign()

	case MedicalHearingAidCompatible:
		// Слуховой аппарат: беспроводная совместимость
		return g.Metadata.WirelessCompatible

	case MedicalWheelchairAccessible:
		// Инвалидная коляска: доступный размер/форма
		return g.Metadata.AccessibleSize ||
			g.IsWheelchairAccessibleByDesign()

	default:
		return true // Неизвестное требование - пропускаем
	}
}

// CompliesWithOther проверяет соответствие подарка прочему требованию.
// Возвращает true если подарок соответствует требованию.
func (g *GiftItem) CompliesWithOther(req OtherRequirement) bool {
	switch req {
	case OtherEcoFriendly:
		// Экологичность: перерабатываемые материалы
		return g.Metadata.EcoFriendly ||
			g.IsMadeFromRecycledMaterials()

	case OtherEducational:
		// Образовательный: развивающий характер
		return g.Metadata.Educational ||
			g.IsEducationalByCategory()

	case OtherGenderNeutral:
		// Гендерно-нейтральный: нейтральные цвета/темы
		return g.Metadata.GenderNeutral ||
			!g.IsGenderSpecific()

	case OtherBilingual:
		// Двуязычный: на двух языках
		return g.Metadata.Bilingual

	case OtherSustainable:
		// Устойчивый: долговечный, ремонтопригодный
		return (g.Metadata.Durable && g.Metadata.Repairable) ||
			g.HasLongWarranty()

	case OtherCharitySupported:
		// Благотворительный: часть средств идет на благотворительность
		return g.Metadata.CharitySupported

	default:
		return true // Неизвестное требование - пропускаем
	}
}

// ==================== Helper методы для проверки дополнительных условий ====================

// ContainsPork проверяет содержит ли продукт свинину.
func (g *GiftItem) ContainsPork() bool {
	// Проверка по названию или материалам
	porkKeywords := []string{"свинина", "pork", "bacon", "ham", "сало", "сальная", "свиной"}
	nameLower := strings.ToLower(g.Name)

	for _, keyword := range porkKeywords {
		if strings.Contains(nameLower, keyword) {
			return true
		}
	}

	// Проверка материалов
	for _, material := range g.Metadata.Materials {
		materialLower := strings.ToLower(material)
		for _, keyword := range porkKeywords {
			if strings.Contains(materialLower, keyword) {
				return true
			}
		}
	}

	// Проверка предупреждений
	for _, warning := range g.Metadata.Warnings {
		warningLower := strings.ToLower(warning)
		for _, keyword := range porkKeywords {
			if strings.Contains(warningLower, keyword) {
				return true
			}
		}
	}

	return false
}

// IsSlaughteredAccordingToHalal проверяет соответствует ли продукт халяль.
func (g *GiftItem) IsSlaughteredAccordingToHalal() bool {
	// Упрощенная проверка - в реальной системе должна быть сложнее
	halalKeywords := []string{"халяль", "halal", "ذَبِيحَة", "халал", "мусульманск", "islamic"}

	// Проверка сертификаций
	for _, cert := range g.Metadata.Certifications {
		certLower := strings.ToLower(cert)
		for _, keyword := range halalKeywords {
			if strings.Contains(certLower, keyword) {
				return true
			}
		}
	}

	// Проверка в названии
	nameLower := strings.ToLower(g.Name)
	for _, keyword := range halalKeywords {
		if strings.Contains(nameLower, keyword) {
			return true
		}
	}

	return false
}

// IsKosherByIngredients проверяет кошерность по ингредиентам.
func (g *GiftItem) IsKosherByIngredients() bool {
	kosherKeywords := []string{"кошер", "kosher", "כָּשֵׁר", "еврейск", "jewish", "иудейск"}

	// Проверка сертификаций
	for _, cert := range g.Metadata.Certifications {
		certLower := strings.ToLower(cert)
		for _, keyword := range kosherKeywords {
			if strings.Contains(certLower, keyword) {
				return true
			}
		}
	}

	// Проверка в названии
	nameLower := strings.ToLower(g.Name)
	for _, keyword := range kosherKeywords {
		if strings.Contains(nameLower, keyword) {
			return true
		}
	}

	// Проверка запрещенных ингредиентов (трефа)
	nonKosherIngredients := []string{
		"свинина", "pork", "моллюски", "shellfish", "ракообразные", "crustaceans",
		"зайчатина", "hare", "верблюжатина", "camel", "хищные птицы", "birds of prey",
		"осетрина", "sturgeon", "сом", "catfish", "угорь", "eel", "акула", "shark",
	}

	// Проверка материалов/ингредиентов на некошерность
	for _, material := range g.Metadata.Materials {
		materialLower := strings.ToLower(material)
		for _, ingredient := range nonKosherIngredients {
			if strings.Contains(materialLower, ingredient) {
				return false
			}
		}
	}

	// Проверка смешивания мяса и молока
	if g.Metadata.ContainsMeat && g.Metadata.ContainsDairy {
		return false
	}

	return true
}

// ContainsAllergenicMaterials проверяет наличие аллергенных материалов.
func (g *GiftItem) ContainsAllergenicMaterials() bool {
	commonAllergens := []string{
		"латекс", "latex", "шерсть", "wool", "пух", "down",
		"пыльца", "pollen", "перо", "feather", "мех", "fur",
		"шелк", "silk", "кашемир", "cashmere", "мохер", "mohair",
		"плюш", "plush", "ворс", "nap", "бархат", "velvet",
	}

	for _, material := range g.Metadata.Materials {
		materialLower := strings.ToLower(material)
		for _, allergen := range commonAllergens {
			if strings.Contains(materialLower, allergen) {
				return true
			}
		}
	}

	return false
}

// HasSafetyCertification проверяет наличие сертификатов безопасности.
func (g *GiftItem) HasSafetyCertification() bool {
	safetyCerts := []string{
		"CE", "EN71", "ASTM", "ISO8124", "СТБ", "ГОСТ", "РСТ",
		"безопасность", "safety", "сертификат", "certificate", "certification",
		"стандарт", "standard", "соответствие", "compliance",
	}

	for _, cert := range g.Metadata.Certifications {
		certUpper := strings.ToUpper(cert)
		for _, safetyCert := range safetyCerts {
			if strings.Contains(certUpper, safetyCert) {
				return true
			}
		}
	}

	// Проверка в предупреждениях
	for _, warning := range g.Metadata.Warnings {
		warningUpper := strings.ToUpper(warning)
		for _, safetyCert := range safetyCerts {
			if strings.Contains(warningUpper, safetyCert) {
				return true
			}
		}
	}

	return false
}

// ContainsBPA проверяет содержит ли продукт BPA.
func (g *GiftItem) ContainsBPA() bool {
	bpaKeywords := []string{"BPA", "бисфенол", "bisphenol", "BISPHENOL"}

	for _, material := range g.Metadata.Materials {
		materialUpper := strings.ToUpper(material)
		for _, keyword := range bpaKeywords {
			if strings.Contains(materialUpper, keyword) {
				return true
			}
		}
	}

	// Проверка в предупреждениях
	for _, warning := range g.Metadata.Warnings {
		warningUpper := strings.ToUpper(warning)
		for _, keyword := range bpaKeywords {
			if strings.Contains(warningUpper, keyword) {
				return true
			}
		}
	}

	return false
}

// IsFocusEnhancing проверяет способствует ли игрушка концентрации.
func (g *GiftItem) IsFocusEnhancing() bool {
	focusCategories := []string{
		"конструктор", "constructor", "пазл", "puzzle", "головоломка",
		"мозаика", "mosaic", "лабиринт", "labyrinth", "сортировщик",
		"логический", "logic", "стратегия", "strategy", "шахматы", "chess",
		"шашки", "checkers", "головолом", "brain teaser",
	}

	categoryLower := strings.ToLower(g.Category)
	nameLower := strings.ToLower(g.Name)

	allText := categoryLower + " " + nameLower
	for _, focusCat := range focusCategories {
		if strings.Contains(allText, focusCat) {
			return true
		}
	}

	return false
}

// IsAutismFriendlyByDesign проверяет дизайн на дружественность к аутизму.
func (g *GiftItem) IsAutismFriendlyByDesign() bool {
	autismFriendlyFeatures := []string{
		"сенсорный", "sensory", "тактильный", "tactile", "успокаивающий",
		"calming", "предсказуемый", "predictable", "структурированный",
		"структурный", "структура", "устойчивый", "stable", "мягкий",
		"soft", "тяжелый", "weighted", "антистресс", "anti-stress",
		"релакс", "relax", "медитатив", "meditative",
	}

	nameLower := strings.ToLower(g.Name)
	categoryLower := strings.ToLower(g.Category)

	allText := nameLower + " " + categoryLower
	for _, feature := range autismFriendlyFeatures {
		if strings.Contains(allText, feature) {
			return true
		}
	}

	return false
}

// IsWheelchairAccessibleByDesign проверяет доступность для инвалидных колясок.
func (g *GiftItem) IsWheelchairAccessibleByDesign() bool {
	// Предметы, которые обычно доступны
	accessibleCategories := []string{
		"книга", "book", "диск", "disc", "программа", "software",
		"аудио", "audio", "видео", "video", "электронный", "electronic",
		"цифровой", "digital", "онлайн", "online", "приложение", "app",
		"музыка", "music", "фильм", "movie", "плеер", "player",
	}

	categoryLower := strings.ToLower(g.Category)
	nameLower := strings.ToLower(g.Name)

	allText := categoryLower + " " + nameLower
	for _, accessibleCat := range accessibleCategories {
		if strings.Contains(allText, accessibleCat) {
			return true
		}
	}

	return false
}

// IsMadeFromRecycledMaterials проверяет сделано ли из переработанных материалов.
func (g *GiftItem) IsMadeFromRecycledMaterials() bool {
	recycledKeywords := []string{
		"переработан", "recycled", "вторичн", "reclaimed", "восстановлен",
		"upcycled", "эко", "eco", "биоразлагаем", "biodegradable",
		"экологичн", "ecological", "природный", "natural", "органическ",
		"organic", "компостируем", "compostable",
	}

	for _, material := range g.Metadata.Materials {
		materialLower := strings.ToLower(material)
		for _, keyword := range recycledKeywords {
			if strings.Contains(materialLower, keyword) {
				return true
			}
		}
	}

	// Проверка в названии
	nameLower := strings.ToLower(g.Name)
	for _, keyword := range recycledKeywords {
		if strings.Contains(nameLower, keyword) {
			return true
		}
	}

	return false
}

// IsEducationalByCategory проверяет образовательную ценность по категории.
func (g *GiftItem) IsEducationalByCategory() bool {
	educationalCategories := []string{
		"обучающий", "educational", "развивающий", "developmental",
		"научный", "scientific", "познавательный", "informative",
		"школьный", "school", "учебный", "study", "лаборатория", "lab",
		"образовательн", "education", "развитие", "development",
		"обучение", "learning", "просвещение", "enlightenment",
		"учебник", "textbook", "атлас", "atlas", "глобус", "globe",
		"химия", "chemistry", "физика", "physics", "биология", "biology",
		"математика", "mathematics", "география", "geography",
	}

	categoryLower := strings.ToLower(g.Category)
	nameLower := strings.ToLower(g.Name)

	allText := categoryLower + " " + nameLower
	for _, eduCat := range educationalCategories {
		if strings.Contains(allText, eduCat) {
			return true
		}
	}

	return false
}

// IsGenderSpecific проверяет гендерную специфичность.
func (g *GiftItem) IsGenderSpecific() bool {
	genderSpecificKeywords := map[string][]string{
		"male": {
			"машинка", "car", "робот", "robot", "солдат", "soldier",
			"пистолет", "gun", "трансформер", "transformer", "супергерой",
			"superhero", "синий", "blue", "техника", "технический",
			"конструктор", "constructor", "космос", "space", "динозавр", "dinosaur",
			"гоночный", "racing", "полицейский", "police", "пожарный", "fire",
			"армия", "army", "танк", "tank", "самолет", "airplane",
		},
		"female": {
			"кукла", "doll", "принцесса", "princess", "пони", "pony",
			"косметика", "cosmetics", "украшение", "jewelry", "розовый",
			"pink", "фея", "fairy", "балет", "ballet", "мода", "fashion",
			"русалка", "mermaid", "единорог", "unicorn", "сердечко", "heart",
			"блеск", "glitter", "пастель", "pastel", "кухня", "kitchen",
			"макияж", "makeup", "платье", "dress", "сумка", "bag",
		},
	}

	nameLower := strings.ToLower(g.Name)
	categoryLower := strings.ToLower(g.Category)

	allText := nameLower + " " + categoryLower

	// Если содержит ключевые слова обоих гендеров - считаем нейтральным
	hasMale := false
	hasFemale := false

	for _, keyword := range genderSpecificKeywords["male"] {
		if strings.Contains(allText, keyword) {
			hasMale = true
			break
		}
	}

	for _, keyword := range genderSpecificKeywords["female"] {
		if strings.Contains(allText, keyword) {
			hasFemale = true
			break
		}
	}

	// Если есть только один тип ключевых слов - предмет гендерно-специфичный
	return (hasMale && !hasFemale) || (hasFemale && !hasMale)
}

// HasLongWarranty проверяет наличие длительной гарантии.
func (g *GiftItem) HasLongWarranty() bool {
	warrantyKeywords := []string{
		"гарантия", "warranty", "гарантийный", "guarantee",
		"пожизненный", "lifetime", "долгий срок", "long term",
		"продленная", "extended", "пятилетняя", "5 year",
		"десятилетняя", "10 year", "пожизненная", "lifelong",
	}

	for _, warning := range g.Metadata.Warnings {
		warningLower := strings.ToLower(warning)
		for _, keyword := range warrantyKeywords {
			if strings.Contains(warningLower, keyword) {
				return true
			}
		}
	}

	// Проверка в названии
	nameLower := strings.ToLower(g.Name)
	for _, keyword := range warrantyKeywords {
		if strings.Contains(nameLower, keyword) {
			return true
		}
	}

	return false
}

// ValidateRequirementsCompliance проверяет соответствие подарка всем требованиям.
// Возвращает список несоответствий или nil если все требования соблюдены.
func (g *GiftItem) ValidateRequirementsCompliance(reqs *SpecialRequirements) []string {
	if reqs == nil {
		return nil
	}

	var violations []string

	// Проверка диетических требований
	for _, dietaryReq := range reqs.Dietary {
		if !g.CompliesWithDietary(dietaryReq) {
			violations = append(violations,
				fmt.Sprintf("Не соответствует диетическому требованию: %s",
					GetDietaryDescription(dietaryReq)))
		}
	}

	// Проверка требований безопасности
	for _, safetyReq := range reqs.Safety {
		if !g.CompliesWithSafety(safetyReq) {
			violations = append(violations,
				fmt.Sprintf("Не соответствует требованию безопасности: %s",
					GetSafetyDescription(safetyReq)))
		}
	}

	// Проверка медицинских требований
	for _, medicalReq := range reqs.Medical {
		if !g.CompliesWithMedical(medicalReq) {
			violations = append(violations,
				fmt.Sprintf("Не соответствует медицинскому требованию: %s",
					GetMedicalDescription(medicalReq)))
		}
	}

	// Проверка прочих требований
	for _, otherReq := range reqs.Other {
		if !g.CompliesWithOther(otherReq) {
			violations = append(violations,
				fmt.Sprintf("Не соответствует прочему требованию: %s",
					GetOtherDescription(otherReq)))
		}
	}

	return violations
}

// CanBeIncludedInGift проверяет может ли предмет быть включен в подарок
// с учетом всех специальных требований ребенка.
// Возвращает true и список предупреждений если предмет подходит.
func (g *GiftItem) CanBeIncludedInGift(child *Child) (bool, []string) {
	if child.SpecialRequirements == nil {
		return g.isAgeAppropriate(child.Age), nil
	}

	violations := g.ValidateRequirementsCompliance(child.SpecialRequirements)
	if len(violations) > 0 {
		return false, violations
	}

	// Проверка возраста
	if !g.isAgeAppropriate(child.Age) {
		return false, []string{fmt.Sprintf(
			"Предмет не подходит по возрасту: требуется %d лет, ребенку %d лет",
			g.MinAge, child.Age)}
	}

	// Проверка особых заметок
	var warnings []string
	if child.Notes != "" {
		// Здесь можно добавить логику анализа заметок
		warnings = append(warnings,
			fmt.Sprintf("У ребенка есть особые заметки: %s", child.Notes))
	}

	return true, warnings
}

// isAgeAppropriate проверяет подходит ли предмет по возрасту.
func (g *GiftItem) isAgeAppropriate(childAge int) bool {
	return childAge >= g.MinAge
}

// GetComplianceSummary возвращает сводку о соответствии требованиям.
func (g *GiftItem) GetComplianceSummary(reqs *SpecialRequirements) map[string]bool {
	summary := make(map[string]bool)

	if reqs == nil {
		return summary
	}

	for _, req := range reqs.Dietary {
		key := fmt.Sprintf("dietary_%s", req)
		summary[key] = g.CompliesWithDietary(req)
	}

	for _, req := range reqs.Safety {
		key := fmt.Sprintf("safety_%s", req)
		summary[key] = g.CompliesWithSafety(req)
	}

	for _, req := range reqs.Medical {
		key := fmt.Sprintf("medical_%s", req)
		summary[key] = g.CompliesWithMedical(req)
	}

	for _, req := range reqs.Other {
		key := fmt.Sprintf("other_%s", req)
		summary[key] = g.CompliesWithOther(req)
	}

	return summary
}

// GetPriceWithCoefficient возвращает цену с учетом коэффициента региона.
func (g *GiftItem) GetPriceWithCoefficient(coefficient float64) float64 {
	if coefficient <= 0 {
		coefficient = 1.0
	}
	return g.Price * coefficient
}

// GetWeightWithCoefficient возвращает вес с учетом коэффициента (если нужно).
func (g *GiftItem) GetWeightWithCoefficient(coefficient float64) float64 {
	if coefficient <= 0 {
		coefficient = 1.0
	}
	return g.Weight * coefficient
}

// ValidateGiftItem проверяет корректность данных предмета подарка.
func ValidateGiftItem(item GiftItem) error {
	if item.ID <= 0 {
		return fmt.Errorf("недопустимый ID предмета: %d", item.ID)
	}

	if strings.TrimSpace(item.Name) == "" {
		return fmt.Errorf("название предмета не может быть пустым")
	}

	if strings.TrimSpace(item.Category) == "" {
		return fmt.Errorf("категория предмета не может быть пустой")
	}

	if item.Price < 0 {
		return fmt.Errorf("цена не может быть отрицательной: %.2f", item.Price)
	}

	if item.Weight < 0 {
		return fmt.Errorf("вес не может быть отрицательным: %.2f", item.Weight)
	}

	if item.MinAge < 0 {
		return fmt.Errorf("минимальный возраст не может быть отрицательным: %d", item.MinAge)
	}

	return nil
}
