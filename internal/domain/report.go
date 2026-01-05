package domain

import (
	"time"
)

// Report представляет полный отчет о расчете подарков.
type Report struct {
	Version          string           `json:"version"`
	GeneratedAt      time.Time        `json:"generated_at"`
	Results          []ChildResult    `json:"results"`
	AgeGroupAnalysis AgeGroupAnalysis `json:"age_group_analysis,omitempty"`
}

// ReportParameters содержит параметры запуска расчета.
type ReportParameters struct {
	ChildrenFile         string  `json:"children_file"`
	CatalogFile          string  `json:"catalog_file"`
	WishesFile           string  `json:"wishes_file,omitempty"`
	RegionsFile          string  `json:"regions_file,omitempty"`
	MaxGiftPrice         float64 `json:"max_gift_price,omitempty"`
	TotalBudget          float64 `json:"total_budget,omitempty"`
	ConsiderRequirements bool    `json:"consider_requirements"`
}

// ReportStatistics содержит статистику расчета.
type ReportStatistics struct {
	TotalChildren          int                    `json:"total_children"`
	SuccessfulCalculations int                    `json:"successful_calculations"`
	FailedCalculations     int                    `json:"failed_calculations"`
	ProcessingTimeMs       int64                  `json:"processing_time_ms"`
	TotalCost              float64                `json:"total_cost"`
	TotalWeight            float64                `json:"total_weight"`
	AverageCostPerChild    float64                `json:"average_cost_per_child"`
	MinGiftCost            float64                `json:"min_gift_cost"`
	MaxGiftCost            float64                `json:"max_gift_cost"`
	AverageItemsPerGift    float64                `json:"average_items_per_gift"`
	BudgetUsagePercentage  float64                `json:"budget_usage_percentage"`
	RequirementsStatistics RequirementsStatistics `json:"requirements_statistics"`
}

// RequirementsStatistics содержит статистику по требованиям.
type RequirementsStatistics struct {
	ChildrenWithDietaryRequirements int    `json:"children_with_dietary_requirements"`
	ChildrenWithSafetyRequirements  int    `json:"children_with_safety_requirements"`
	ChildrenWithMedicalRequirements int    `json:"children_with_medical_requirements"`
	ChildrenWithOtherRequirements   int    `json:"children_with_other_requirements"`
	MostCommonDietary               string `json:"most_common_dietary,omitempty"`
	MostCommonSafety                string `json:"most_common_safety,omitempty"`
}

// ChildResult содержит результат расчета для одного ребенка.
type ChildResult struct {
	ChildID             int                  `json:"child_id"`
	ChildName           string               `json:"child_name"`
	Age                 int                  `json:"age"`
	Region              string               `json:"region"`
	SpecialRequirements *SpecialRequirements `json:"special_requirements,omitempty"`
	GiftSelection       []GiftSelection      `json:"gift_selection"`
	CostSummary         ChildCostSummary     `json:"cost_summary"`
	SelectionNotes      []string             `json:"selection_notes,omitempty"`
	Warnings            []string             `json:"warnings,omitempty"`
	Errors              *string              `json:"errors,omitempty"`
}

// GiftSelection содержит информацию о выбранном предмете.
type GiftSelection struct {
	ItemID          int                    `json:"item_id"`
	ItemName        string                 `json:"item_name"`
	Category        string                 `json:"category"`
	Price           float64                `json:"price"`
	Weight          float64                `json:"weight"`
	SelectionReason string                 `json:"selection_reason"`
	ComplianceCheck map[string]bool        `json:"compliance_check"`
	MetadataSummary map[string]interface{} `json:"metadata_summary,omitempty"`
}

// ChildCostSummary содержит сводку по стоимости подарка.
type ChildCostSummary struct {
	Cost       float64 `json:"cost"`
	ItemsCount int     `json:"items_count"`
}

// FailedCalculation содержит информацию о неудачном расчете.
type FailedCalculation struct {
	ChildID              int                   `json:"child_id"`
	ChildName            string                `json:"child_name"`
	Age                  int                   `json:"age"`
	Region               string                `json:"region"`
	ErrorType            string                `json:"error_type"`
	ErrorMessage         string                `json:"error_message"`
	RequirementsConflict *RequirementsConflict `json:"requirements_conflict,omitempty"`
	PartialSelection     []GiftSelection       `json:"partial_selection,omitempty"`
	Suggestions          []string              `json:"suggestions,omitempty"`
}

// RequirementsConflict содержит информацию о конфликте требований.
type RequirementsConflict struct {
	Dietary         []DietaryRequirement `json:"dietary,omitempty"`
	FailedItems     []int                `json:"failed_items,omitempty"`
	ConflictDetails string               `json:"conflict_details,omitempty"`
}

// ProductionSummary содержит производственную сводку.
type ProductionSummary struct {
	TotalItemsNeeded    int                           `json:"total_items_needed"`
	ItemsBreakdown      []ProductionItemBreakdown     `json:"items_breakdown"`
	CategoriesBreakdown []ProductionCategoryBreakdown `json:"categories_breakdown"`
}

// ProductionItemBreakdown содержит разбивку по предметам.
type ProductionItemBreakdown struct {
	ItemID           int     `json:"item_id"`
	ItemName         string  `json:"item_name"`
	Category         string  `json:"category"`
	RequiredQuantity int     `json:"required_quantity"`
	TotalCost        float64 `json:"total_cost"`
	TotalWeight      float64 `json:"total_weight"`
}

// ProductionCategoryBreakdown содержит разбивку по категориям.
type ProductionCategoryBreakdown struct {
	CategoryID    string  `json:"category_id"`
	CategoryName  string  `json:"category_name"`
	ItemsCount    int     `json:"items_count"`
	TotalQuantity int     `json:"total_quantity"`
	TotalCost     float64 `json:"total_cost"`
}

// BudgetAnalysis содержит анализ бюджета.
type BudgetAnalysis struct {
	TotalBudget     float64  `json:"total_budget"`
	TotalUsed       float64  `json:"total_used"`
	RemainingBudget float64  `json:"remaining_budget"`
	UsagePercentage float64  `json:"usage_percentage"`
	PerChildAverage float64  `json:"per_child_average"`
	BudgetStatus    string   `json:"budget_status"` // UNDER_BUDGET, WITHIN_BUDGET, OVER_BUDGET
	Recommendations []string `json:"recommendations"`
}

// RegionAnalysis содержит анализ по регионам.
type RegionAnalysis struct {
	Region        string  `json:"region"`
	ChildrenCount int     `json:"children_count"`
	TotalCost     float64 `json:"total_cost"`
	AverageCost   float64 `json:"average_cost"`
	Coefficient   float64 `json:"coefficient"`
}

// AgeGroupAnalysis содержит анализ по возрастным группам.
type AgeGroupAnalysis struct {
	AgeGroup      string  `json:"age_group"`
	MinAge        int     `json:"min_age"`
	MaxAge        int     `json:"max_age"`
	ChildrenCount int     `json:"children_count"`
	TotalCost     float64 `json:"total_cost"`
	AverageCost   float64 `json:"average_cost"`
}

// RequirementsAnalysis содержит анализ влияния требований.
type RequirementsAnalysis struct {
	DietaryImpact map[string]RequirementImpact `json:"dietary_impact,omitempty"`
	SafetyImpact  map[string]RequirementImpact `json:"safety_impact,omitempty"`
	MedicalImpact map[string]RequirementImpact `json:"medical_impact,omitempty"`
	OtherImpact   map[string]RequirementImpact `json:"other_impact,omitempty"`
}

// RequirementImpact содержит информацию о влиянии требования.
type RequirementImpact struct {
	AffectedChildren    int     `json:"affected_children"`
	AverageCostIncrease float64 `json:"average_cost_increase"`
	DifficultyLevel     string  `json:"difficulty_level"` // LOW, MEDIUM, HIGH
}

// OptimizationSuggestion содержит предложение по оптимизации.
type OptimizationSuggestion struct {
	Suggestion       string   `json:"suggestion"`
	PotentialSaving  float64  `json:"potential_saving"`
	AffectedChildren []int    `json:"affected_children,omitempty"`
	AffectedRegions  []string `json:"affected_regions,omitempty"`
	Complexity       string   `json:"complexity"` // LOW, MEDIUM, HIGH
}
