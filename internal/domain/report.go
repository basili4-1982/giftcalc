package domain

import "time"

type GiftSelection struct {
	Item   GiftItem `json:"item"`
	Reason string   `json:"reason"` // "wish", "age_appropriate", "fallback"
}

type ChildResult struct {
	ChildID     int             `json:"child_id"`
	ChildName   string          `json:"child_name,omitempty"`
	Selections  []GiftSelection `json:"selections"`
	TotalCost   float64         `json:"total_cost"`
	TotalWeight float64         `json:"total_weight"`
	Region      string          `json:"region"`
	Coefficient float64         `json:"coefficient"`
	FinalCost   float64         `json:"final_cost"`
	Error       string          `json:"error,omitempty"`
}

type Report struct {
	GeneratedAt time.Time     `json:"generated_at"`
	Version     string        `json:"version"`
	Parameters  ReportParams  `json:"parameters"`
	Results     []ChildResult `json:"results"`
	Summary     ReportSummary `json:"summary"`
	TotalCost   float64       `json:"total_cost"`
}

type ReportParams struct {
	MaxGiftPrice float64 `json:"max_gift_price,omitempty"`
	TotalBudget  float64 `json:"total_budget,omitempty"`
	DataFiles    struct {
		Children string `json:"children,omitempty"`
		Catalog  string `json:"catalog,omitempty"`
		Wishes   string `json:"wishes,omitempty"`
		Regions  string `json:"regions,omitempty"`
	} `json:"data_files"`
}

type ReportSummary struct {
	TotalChildren  int     `json:"total_children"`
	Successful     int     `json:"successful"`
	Failed         int     `json:"failed"`
	AvgCostPerGift float64 `json:"avg_cost_per_gift"`
	MinCost        float64 `json:"min_cost"`
	MaxCost        float64 `json:"max_cost"`
	TotalWeight    float64 `json:"total_weight"`
	BudgetUsage    float64 `json:"budget_usage,omitempty"`
}
