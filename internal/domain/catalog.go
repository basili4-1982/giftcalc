package domain

import "time"

type CatalogItem struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Price    float64  `json:"price"`
	Weight   float64  `json:"weight"`
	MinAge   int      `json:"min_age"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	ContainsNuts       bool     `json:"contains_nuts,omitempty"`
	ContainsSugar      bool     `json:"contains_sugar,omitempty"`
	ContainsDairy      bool     `json:"contains_dairy,omitempty"`
	ContainsGluten     bool     `json:"contains_gluten,omitempty"`
	Vegetarian         bool     `json:"vegetarian,omitempty"`
	Vegan              bool     `json:"vegan,omitempty"`
	SugarFree          bool     `json:"sugar_free,omitempty"`
	HalalCertified     bool     `json:"halal_certified,omitempty"`
	KosherCertified    bool     `json:"kosher_certified,omitempty"`
	Hypoallergenic     bool     `json:"hypoallergenic"`
	NonToxic           bool     `json:"non_toxic"`
	EcoFriendly        bool     `json:"eco_friendly"`
	Educational        bool     `json:"educational"`
	GenderNeutral      bool     `json:"gender_neutral"`
	Materials          []string `json:"materials"`
	Certifications     []string `json:"certifications"`
	Warnings           []string `json:"warnings"`
	HasSmallParts      bool     `json:"has_small_parts,omitempty"`
	SmallPartsSize     float64  `json:"small_parts_size,omitempty"`
	Washable           bool     `json:"washable,omitempty"`
	FlameRetardant     bool     `json:"flame_retardant,omitempty"`
	BpaFree            bool     `json:"bpa_free,omitempty"`
	HasFlashingLights  bool     `json:"has_flashing_lights,omitempty"`
	HasFuzzyMaterial   bool     `json:"has_fuzzy_material,omitempty"`
	IsDusty            bool     `json:"is_dusty,omitempty"`
	CalmingEffect      bool     `json:"calming_effect,omitempty"`
	Tactile            bool     `json:"tactile,omitempty"`
	Predictable        bool     `json:"predictable,omitempty"`
	Durable            bool     `json:"durable,omitempty"`
	Repairable         bool     `json:"repairable,omitempty"`
	Bilingual          bool     `json:"bilingual,omitempty"`
	AccessibleSize     bool     `json:"accessible_size,omitempty"`
	CharitySupported   bool     `json:"charity_supported,omitempty"`
	WirelessCompatible bool     `json:"wireless_compatible,omitempty"`
}

type CatalogData struct {
	Version     string    `json:"version"`
	GeneratedAt time.Time `json:"generated_at"`
	Description string    `json:"description"`
	Metadata    struct {
		TotalItems      int `json:"total_items"`
		TotalCategories int `json:"total_categories"`
		PriceRange      struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
			Avg float64 `json:"avg"`
		} `json:"price_range"`
		AgeRange struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"age_range"`
		LastUpdated string `json:"last_updated"`
		Source      string `json:"source"`
	} `json:"metadata"`
	Categories []struct {
		Id          string  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		BaseWeight  float64 `json:"base_weight"`
		MinAge      int     `json:"min_age"`
		MaxAge      int     `json:"max_age"`
	} `json:"categories"`
	Items []CatalogItem `json:"items"`
}
