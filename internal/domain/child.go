package domain

import (
	"fmt"
	"strings"
	"time"
)

// DietaryRequirement представляет диетические ограничения и предпочтения ребенка.
// Эти требования влияют на выбор продуктов питания в подарках.
type DietaryRequirement string

const (
	// DietaryVegetarian - вегетарианство.
	// Влияние на подарок: исключить продукты с мясом, рыбой
	DietaryVegetarian DietaryRequirement = "vegetarian"

	// DietaryVegan - веганство.
	// Влияние на подарок: исключить все продукты животного происхождения
	DietaryVegan DietaryRequirement = "vegan"

	// DietaryNutsAllergy - аллергия на орехи.
	// Влияние на подарок: исключить орехи и следы орехов
	DietaryNutsAllergy DietaryRequirement = "nuts_allergy"

	// DietaryLactoseIntolerant - непереносимость лактозы.
	// Влияние на подарок: исключить молочные продукты
	DietaryLactoseIntolerant DietaryRequirement = "lactose_intolerant"

	// DietaryGlutenFree - без глютена.
	// Влияние на подарок: исключить пшеницу, ячмень, рожь
	DietaryGlutenFree DietaryRequirement = "gluten_free"

	// DietaryDiabetes - диабет.
	// Влияние на подарок: ограничить сахар, использовать сахарозаменители
	DietaryDiabetes DietaryRequirement = "diabetes"

	// DietaryHalal - халяль.
	// Влияние на подарок: соответствие исламским нормам
	DietaryHalal DietaryRequirement = "halal"

	// DietaryKosher - кошерность.
	// Влияние на подарок: соответствие иудейским нормам
	DietaryKosher DietaryRequirement = "kosher"
)

// SafetyRequirement представляет требования безопасности для игрушек и подарков.
// Эти требования особенно важны для детей младшего возраста.
type SafetyRequirement string

const (
	// SafetyNoSmallParts - без мелких деталей.
	// Влияние на подарок: детали > 3см для детей до 3 лет
	SafetyNoSmallParts SafetyRequirement = "no_small_parts"

	// SafetyHypoallergenic - гипоаллергенность.
	// Влияние на подарок: без аллергенов в материалах
	SafetyHypoallergenic SafetyRequirement = "hypoallergenic"

	// SafetyNonToxic - нетоксичность.
	// Влияние на подарок: сертифицированные безопасные материалы
	SafetyNonToxic SafetyRequirement = "non_toxic"

	// SafetyWashable - моющийся.
	// Влияние на подарок: возможность стирки/мытья
	SafetyWashable SafetyRequirement = "washable"

	// SafetyFlameRetardant - огнестойкость.
	// Влияние на подарок: трудновоспламеняемые материалы
	SafetyFlameRetardant SafetyRequirement = "flame_retardant"

	// SafetyBPAFree - без BPA.
	// Влияние на подарок: без бисфенола-А в пластике
	SafetyBPAFree SafetyRequirement = "bpa_free"
)

// MedicalRequirement представляет медицинские особенности ребенка.
// Эти требования учитывают специфические медицинские условия.
type MedicalRequirement string

const (
	// MedicalEpilepsy - эпилепсия.
	// Влияние на подарок: избегать мигающих светодиодов
	MedicalEpilepsy MedicalRequirement = "epilepsy"

	// MedicalAsthma - астма.
	// Влияние на подарок: избегать пушистых/пылящих материалов
	MedicalAsthma MedicalRequirement = "asthma"

	// MedicalADHDFriendly - для СДВГ.
	// Влияние на подарок: успокаивающие, фокусирующие игрушки
	MedicalADHDFriendly MedicalRequirement = "adhd_friendly"

	// MedicalAutismFriendly - для аутизма.
	// Влияние на подарок: тактильные, предсказуемые игрушки
	MedicalAutismFriendly MedicalRequirement = "autism_friendly"

	// MedicalHearingAidCompatible - слуховой аппарат.
	// Влияние на подарок: беспроводная совместимость
	MedicalHearingAidCompatible MedicalRequirement = "hearing_aid_compatible"

	// MedicalWheelchairAccessible - инвалидная коляска.
	// Влияние на подарок: доступный размер/форма
	MedicalWheelchairAccessible MedicalRequirement = "wheelchair_accessible"
)

// OtherRequirement представляет прочие требования и предпочтения.
// Эти требования отражают этические, образовательные и социальные аспекты.
type OtherRequirement string

const (
	// OtherEcoFriendly - экологичность.
	// Влияние на подарок: перерабатываемые материалы
	OtherEcoFriendly OtherRequirement = "eco_friendly"

	// OtherEducational - образовательный.
	// Влияние на подарок: развивающий характер
	OtherEducational OtherRequirement = "educational"

	// OtherGenderNeutral - гендерно-нейтральный.
	// Влияние на подарок: нейтральные цвета/темы
	OtherGenderNeutral OtherRequirement = "gender_neutral"

	// OtherBilingual - двуязычный.
	// Влияние на подарок: на двух языках
	OtherBilingual OtherRequirement = "bilingual"

	// OtherSustainable - устойчивый.
	// Влияние на подарок: долговечный, ремонтопригодный
	OtherSustainable OtherRequirement = "sustainable"

	// OtherCharitySupported - благотворительный.
	// Влияние на подарок: часть средств идет на благотворительность
	OtherCharitySupported OtherRequirement = "charity_supported"
)

// SpecialRequirements представляет все специальные требования ребенка.
// Структура организована по категориям для удобства обработки и валидации.
type SpecialRequirements struct {
	// Dietary содержит диетические ограничения ребенка.
	// Может быть nil или пустым массивом если ограничений нет.
	Dietary []DietaryRequirement `json:"dietary,omitempty"`

	// Safety содержит требования безопасности для игрушек.
	// Особенно важно для детей младшего возраста.
	Safety []SafetyRequirement `json:"safety,omitempty"`

	// Medical содержит медицинские особенности ребенка.
	// Учитывает специфические медицинские условия.
	Medical []MedicalRequirement `json:"medical,omitempty"`

	// Other содержит прочие требования и предпочтения.
	// Отражает этические, образовательные и социальные аспекты.
	Other []OtherRequirement `json:"other,omitempty"`
}

type ChildrenData struct {
	Version     string    `json:"version"`
	GeneratedAt time.Time `json:"generated_at"`
	Description string    `json:"description"`
	Metadata    struct {
		TotalCount int      `json:"total_count"`
		Regions    []string `json:"regions"`
		AgeRange   struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"age_range"`
		Source                  string `json:"source"`
		DataFormat              string `json:"data_format"`
		SpecialRequirementsKeys struct {
			Dietary []string `json:"dietary"`
			Safety  []string `json:"safety"`
			Medical []string `json:"medical"`
			Other   []string `json:"other"`
		} `json:"special_requirements_keys"`
	} `json:"metadata"`
	Children []Child `json:"children"`
}

// Child представляет информацию о ребенке.
// Используется для расчета индивидуального подарка.
type Child struct {
	// ID - уникальный идентификатор ребенка.
	// Должен быть уникальным в пределах системы.
	ID int `json:"id"`

	// Name - имя ребенка.
	// Может содержать кириллицу и пробелы.
	Name string `json:"name"`

	// Age - возраст ребенка в полных годах.
	// Используется для возрастной фильтрации подарков.
	Age int `json:"age"`

	// Region - регион проживания ребенка.
	// Используется для расчета региональных коэффициентов.
	Region string `json:"region"`

	// Notes - дополнительные заметки о ребенке.
	// Необязательное поле, может содержать произвольный текст.
	Notes string `json:"notes,omitempty"`

	// Tags - теги для категоризации детей.
	// Используется для групповой обработки и аналитики.
	Tags []string `json:"tags,omitempty"`

	// SpecialRequirements - специальные требования ребенка.
	// Учитываются при подборе подарков для обеспечения безопасности и комфорта.
	SpecialRequirements *SpecialRequirements `json:"special_requirements,omitempty"`
}

// ChildRepository определяет интерфейс для работы с данными о детях.
type ChildRepository interface {
	// GetAll возвращает всех детей.
	GetAll() ([]Child, error)

	// GetByID возвращает ребенка по идентификатору.
	GetByID(id int) (*Child, error)

	// GetByRegion возвращает детей по региону.
	GetByRegion(region string) ([]Child, error)

	// GetByAgeRange возвращает детей в указанном возрастном диапазоне.
	GetByAgeRange(minAge, maxAge int) ([]Child, error)

	// GetByTags возвращает детей с указанными тегами.
	GetByTags(tags []string) ([]Child, error)
}

// Validate проверяет корректность специальных требований.
// Возвращает ошибку если найдены недопустимые значения.
func (sr *SpecialRequirements) Validate() error {
	if sr == nil {
		return nil
	}

	// Валидация диетических требований
	validDietary := map[DietaryRequirement]bool{
		DietaryVegetarian:        true,
		DietaryVegan:             true,
		DietaryNutsAllergy:       true,
		DietaryLactoseIntolerant: true,
		DietaryGlutenFree:        true,
		DietaryDiabetes:          true,
		DietaryHalal:             true,
		DietaryKosher:            true,
	}

	for _, req := range sr.Dietary {
		if !validDietary[req] {
			return fmt.Errorf("недопустимое диетическое требование: %s", req)
		}
	}

	// Валидация требований безопасности
	validSafety := map[SafetyRequirement]bool{
		SafetyNoSmallParts:   true,
		SafetyHypoallergenic: true,
		SafetyNonToxic:       true,
		SafetyWashable:       true,
		SafetyFlameRetardant: true,
		SafetyBPAFree:        true,
	}

	for _, req := range sr.Safety {
		if !validSafety[req] {
			return fmt.Errorf("недопустимое требование безопасности: %s", req)
		}
	}

	// Валидация медицинских требований
	validMedical := map[MedicalRequirement]bool{
		MedicalEpilepsy:             true,
		MedicalAsthma:               true,
		MedicalADHDFriendly:         true,
		MedicalAutismFriendly:       true,
		MedicalHearingAidCompatible: true,
		MedicalWheelchairAccessible: true,
	}

	for _, req := range sr.Medical {
		if !validMedical[req] {
			return fmt.Errorf("недопустимое медицинское требование: %s", req)
		}
	}

	// Валидация прочих требований
	validOther := map[OtherRequirement]bool{
		OtherEcoFriendly:      true,
		OtherEducational:      true,
		OtherGenderNeutral:    true,
		OtherBilingual:        true,
		OtherSustainable:      true,
		OtherCharitySupported: true,
	}

	for _, req := range sr.Other {
		if !validOther[req] {
			return fmt.Errorf("недопустимое прочее требование: %s", req)
		}
	}

	return nil
}

// HasRequirement проверяет наличие конкретного требования.
// Удобно для фильтрации и проверок.
func (sr *SpecialRequirements) HasRequirement(category, value string) bool {
	if sr == nil {
		return false
	}

	switch category {
	case "dietary":
		for _, req := range sr.Dietary {
			if string(req) == value {
				return true
			}
		}
	case "safety":
		for _, req := range sr.Safety {
			if string(req) == value {
				return true
			}
		}
	case "medical":
		for _, req := range sr.Medical {
			if string(req) == value {
				return true
			}
		}
	case "other":
		for _, req := range sr.Other {
			if string(req) == value {
				return true
			}
		}
	}

	return false
}

// GetRequirementsCount возвращает количество требований по категориям.
// Полезно для статистики и аналитики.
func (sr *SpecialRequirements) GetRequirementsCount() map[string]int {
	result := make(map[string]int)

	if sr == nil {
		return result
	}

	result["dietary"] = len(sr.Dietary)
	result["safety"] = len(sr.Safety)
	result["medical"] = len(sr.Medical)
	result["other"] = len(sr.Other)
	result["total"] = len(sr.Dietary) + len(sr.Safety) + len(sr.Medical) + len(sr.Other)

	return result
}

// String возвращает строковое представление требований.
// Удобно для логгирования и отладки.
func (sr *SpecialRequirements) String() string {
	if sr == nil {
		return "нет требований"
	}

	var parts []string

	if len(sr.Dietary) > 0 {
		parts = append(parts, fmt.Sprintf("диетические: %v", sr.Dietary))
	}
	if len(sr.Safety) > 0 {
		parts = append(parts, fmt.Sprintf("безопасность: %v", sr.Safety))
	}
	if len(sr.Medical) > 0 {
		parts = append(parts, fmt.Sprintf("медицинские: %v", sr.Medical))
	}
	if len(sr.Other) > 0 {
		parts = append(parts, fmt.Sprintf("прочие: %v", sr.Other))
	}

	if len(parts) == 0 {
		return "нет требований"
	}

	return strings.Join(parts, ", ")
}

// GetDietaryDescription возвращает описание диетического требования.
func GetDietaryDescription(req DietaryRequirement) string {
	switch req {
	case DietaryVegetarian:
		return "Вегетарианство - исключить продукты с мясом, рыбой"
	case DietaryVegan:
		return "Веганство - исключить все продукты животного происхождения"
	case DietaryNutsAllergy:
		return "Аллергия на орехи - исключить орехи и следы орехов"
	case DietaryLactoseIntolerant:
		return "Непереносимость лактозы - исключить молочные продукты"
	case DietaryGlutenFree:
		return "Без глютена - исключить пшеницу, ячмень, рожь"
	case DietaryDiabetes:
		return "Диабет - ограничить сахар, использовать сахарозаменители"
	case DietaryHalal:
		return "Халяль - соответствие исламским нормам"
	case DietaryKosher:
		return "Кошерность - соответствие иудейским нормам"
	default:
		return "Неизвестное диетическое требование"
	}
}

// GetSafetyDescription возвращает описание требования безопасности.
func GetSafetyDescription(req SafetyRequirement) string {
	switch req {
	case SafetyNoSmallParts:
		return "Без мелких деталей - детали > 3см для детей до 3 лет"
	case SafetyHypoallergenic:
		return "Гипоаллергенность - без аллергенов в материалах"
	case SafetyNonToxic:
		return "Нетоксичность - сертифицированные безопасные материалы"
	case SafetyWashable:
		return "Моющийся - возможность стирки/мытья"
	case SafetyFlameRetardant:
		return "Огнестойкость - трудновоспламеняемые материалы"
	case SafetyBPAFree:
		return "Без BPA - без бисфенола-А в пластике"
	default:
		return "Неизвестное требование безопасности"
	}
}

// GetMedicalDescription возвращает описание медицинского требования.
func GetMedicalDescription(req MedicalRequirement) string {
	switch req {
	case MedicalEpilepsy:
		return "Эпилепсия - избегать мигающих светодиодов"
	case MedicalAsthma:
		return "Астма - избегать пушистых/пылящих материалов"
	case MedicalADHDFriendly:
		return "Для СДВГ - успокаивающие, фокусирующие игрушки"
	case MedicalAutismFriendly:
		return "Для аутизма - тактильные, предсказуемые игрушки"
	case MedicalHearingAidCompatible:
		return "Слуховой аппарат - беспроводная совместимость"
	case MedicalWheelchairAccessible:
		return "Инвалидная коляска - доступный размер/форма"
	default:
		return "Неизвестное медицинское требование"
	}
}

// GetOtherDescription возвращает описание прочего требования.
func GetOtherDescription(req OtherRequirement) string {
	switch req {
	case OtherEcoFriendly:
		return "Экологичность - перерабатываемые материалы"
	case OtherEducational:
		return "Образовательный - развивающий характер"
	case OtherGenderNeutral:
		return "Гендерно-нейтральный - нейтральные цвета/темы"
	case OtherBilingual:
		return "Двуязычный - на двух языках"
	case OtherSustainable:
		return "Устойчивый - долговечный, ремонтопригодный"
	case OtherCharitySupported:
		return "Благотворительный - часть средств идет на благотворительность"
	default:
		return "Неизвестное прочее требование"
	}
}

// GetAllRequirements возвращает карту всех допустимых требований по категориям.
// Полезно для валидации и генерации документации.
func GetAllRequirements() map[string][]string {
	return map[string][]string{
		"dietary": {
			string(DietaryVegetarian),
			string(DietaryVegan),
			string(DietaryNutsAllergy),
			string(DietaryLactoseIntolerant),
			string(DietaryGlutenFree),
			string(DietaryDiabetes),
			string(DietaryHalal),
			string(DietaryKosher),
		},
		"safety": {
			string(SafetyNoSmallParts),
			string(SafetyHypoallergenic),
			string(SafetyNonToxic),
			string(SafetyWashable),
			string(SafetyFlameRetardant),
			string(SafetyBPAFree),
		},
		"medical": {
			string(MedicalEpilepsy),
			string(MedicalAsthma),
			string(MedicalADHDFriendly),
			string(MedicalAutismFriendly),
			string(MedicalHearingAidCompatible),
			string(MedicalWheelchairAccessible),
		},
		"other": {
			string(OtherEcoFriendly),
			string(OtherEducational),
			string(OtherGenderNeutral),
			string(OtherBilingual),
			string(OtherSustainable),
			string(OtherCharitySupported),
		},
	}
}

// ValidateChild проверяет корректность данных ребенка.
func ValidateChild(child Child) error {
	if child.ID <= 0 {
		return fmt.Errorf("недопустимый ID ребенка: %d", child.ID)
	}

	if strings.TrimSpace(child.Name) == "" {
		return fmt.Errorf("имя ребенка не может быть пустым")
	}

	if child.Age < 0 || child.Age > 18 {
		return fmt.Errorf("недопустимый возраст: %d", child.Age)
	}

	if strings.TrimSpace(child.Region) == "" {
		return fmt.Errorf("регион не может быть пустым")
	}

	// Валидация специальных требований
	if child.SpecialRequirements != nil {
		if err := child.SpecialRequirements.Validate(); err != nil {
			return fmt.Errorf("ошибка валидации специальных требований: %w", err)
		}
	}

	return nil
}

// AgeGroup возвращает возрастную группу ребенка.
func (c *Child) AgeGroup() string {
	switch {
	case c.Age < 4:
		return "toddlers" // Малыши (0-3)
	case c.Age < 7:
		return "preschoolers" // Дошкольники (4-6)
	case c.Age < 11:
		return "young_school" // Младшие школьники (7-10)
	case c.Age < 15:
		return "teens" // Подростки (11-14)
	default:
		return "older_teens" // Старшие подростки (15+)
	}
}

// GetDietaryRequirements возвращает диетические требования как строки.
func (sr *SpecialRequirements) GetDietaryRequirements() []string {
	if sr == nil {
		return nil
	}

	result := make([]string, len(sr.Dietary))
	for i, req := range sr.Dietary {
		result[i] = string(req)
	}
	return result
}

// GetSafetyRequirements возвращает требования безопасности как строки.
func (sr *SpecialRequirements) GetSafetyRequirements() []string {
	if sr == nil {
		return nil
	}

	result := make([]string, len(sr.Safety))
	for i, req := range sr.Safety {
		result[i] = string(req)
	}
	return result
}

// GetMedicalRequirements возвращает медицинские требования как строки.
func (sr *SpecialRequirements) GetMedicalRequirements() []string {
	if sr == nil {
		return nil
	}

	result := make([]string, len(sr.Medical))
	for i, req := range sr.Medical {
		result[i] = string(req)
	}
	return result
}

// GetOtherRequirements возвращает прочие требования как строки.
func (sr *SpecialRequirements) GetOtherRequirements() []string {
	if sr == nil {
		return nil
	}

	result := make([]string, len(sr.Other))
	for i, req := range sr.Other {
		result[i] = string(req)
	}
	return result
}

// HasAnyRequirements проверяет есть ли у ребенка какие-либо специальные требования.
func (c *Child) HasAnyRequirements() bool {
	if c.SpecialRequirements == nil {
		return false
	}

	return len(c.SpecialRequirements.Dietary) > 0 ||
		len(c.SpecialRequirements.Safety) > 0 ||
		len(c.SpecialRequirements.Medical) > 0 ||
		len(c.SpecialRequirements.Other) > 0
}

// RequirementsSummary возвращает краткое описание требований ребенка.
func (c *Child) RequirementsSummary() string {
	if !c.HasAnyRequirements() {
		return "нет особых требований"
	}

	var summary []string

	if c.SpecialRequirements != nil {
		if len(c.SpecialRequirements.Dietary) > 0 {
			summary = append(summary, fmt.Sprintf("диета: %d", len(c.SpecialRequirements.Dietary)))
		}
		if len(c.SpecialRequirements.Safety) > 0 {
			summary = append(summary, fmt.Sprintf("безопасность: %d", len(c.SpecialRequirements.Safety)))
		}
		if len(c.SpecialRequirements.Medical) > 0 {
			summary = append(summary, fmt.Sprintf("медицина: %d", len(c.SpecialRequirements.Medical)))
		}
		if len(c.SpecialRequirements.Other) > 0 {
			summary = append(summary, fmt.Sprintf("прочее: %d", len(c.SpecialRequirements.Other)))
		}
	}

	return strings.Join(summary, ", ")
}

// ChildJSON представляет структуру JSON файла с детьми.
type ChildJSON struct {
	Version     string        `json:"version,omitempty"`
	GeneratedAt string        `json:"generated_at,omitempty"`
	Description string        `json:"description,omitempty"`
	Metadata    ChildMetadata `json:"metadata,omitempty"`
	Children    []Child       `json:"children"`
}

// ChildMetadata содержит метаданные файла с детьми.
type ChildMetadata struct {
	TotalCount   int             `json:"total_count,omitempty"`
	Regions      []string        `json:"regions,omitempty"`
	AgeRange     AgeRange        `json:"age_range,omitempty"`
	Source       string          `json:"source,omitempty"`
	DataFormat   string          `json:"data_format,omitempty"`
	Requirements RequirementKeys `json:"special_requirements_keys,omitempty"`
}

// AgeRange представляет диапазон возрастов.
type AgeRange struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// RequirementKeys содержит ключи специальных требований.
type RequirementKeys struct {
	Dietary []string `json:"dietary,omitempty"`
	Safety  []string `json:"safety,omitempty"`
	Medical []string `json:"medical,omitempty"`
	Other   []string `json:"other,omitempty"`
}
