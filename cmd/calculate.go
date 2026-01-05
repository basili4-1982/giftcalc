package main

import (
	"encoding/json"
	"giftcalc/internal/domain"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var calculateCmd = &cobra.Command{
	Use:   "calculate",
	Short: "Рассчитать стоимость и комплектацию подарков",
	Run:   runCalculate,
}

func init() {
	calculateCmd.
		Flags().
		String("children", "", "Файл с данными о детях (JSON)")
	calculateCmd.
		Flags().String("catalog", "", "Файл каталога подарков")
	calculateCmd.
		Flags().String("report", "report.json", "Файл отчета")
	calculateCmd.
		Flags().Float32("maxBudget", 100, "Максимальный бюджет")
	calculateCmd.
		Flags().Int32("maxCount", 10, "Максимальное количество позиций")
}

func runCalculate(cmd *cobra.Command, args []string) {
	childrenDataFile, err := cmd.Flags().GetString("children")
	if err != nil {
		return
	}

	if childrenDataFile == "" {
		slog.Error("Необходимо передать данные о детях")
		return
	}

	catalogDataFile, err := cmd.Flags().GetString("children")
	if err != nil {
		return
	}

	reportFile, err := cmd.Flags().GetString("report")
	if err != nil {
		return
	}

	if catalogDataFile == "" {
		slog.Error("Необходимо передать данные каталога")
		return
	}

	childrenBinData, err := os.ReadFile(childrenDataFile)
	if err != nil {
		slog.Error("Не могу найти файл '" + childrenDataFile + "'")
		return
	}

	catalogBinData, err := os.ReadFile(catalogDataFile)
	if err != nil {
		slog.Error("Не могу найти файл '" + childrenDataFile + "'")
		return
	}

	childrenData := domain.ChildrenData{}

	err = json.Unmarshal(childrenBinData, &childrenData)
	if err != nil {
		slog.Error("Не могу разобрать файл '"+childrenDataFile+"'", slog.String("err", err.Error()))
		return
	}

	catalog := domain.CatalogData{}
	err = json.Unmarshal(catalogBinData, &catalog)

	report := domain.Report{
		Version:     "v1.0.0",
		GeneratedAt: time.Now(),
	}

	report.Results = make([]domain.ChildResult, 0, len(childrenData.Children))

	for _, child := range childrenData.Children {
		giftSelection := make([]domain.GiftSelection, 0)
		cost := 0.0
		//TODO: написать код подбора подарка

		report.Results = append(report.Results, domain.ChildResult{
			ChildID:             child.ID,
			ChildName:           child.Name,
			Age:                 child.Age,
			Region:              child.Region,
			SpecialRequirements: child.SpecialRequirements,
			GiftSelection:       giftSelection,
			CostSummary: domain.ChildCostSummary{
				Cost:       cost,
				ItemsCount: len(giftSelection),
			},
		})
	}

	report.AgeGroupAnalysis = domain.AgeGroupAnalysis{
		AgeGroup:      "",
		MinAge:        0,
		MaxAge:        0,
		ChildrenCount: len(childrenData.Children),
		TotalCost:     0,
		AverageCost:   0,
	}

	data, err := json.Marshal(report)
	if err != nil {
		slog.Error("Не смог сформировать файл отчета", slog.String("err", err.Error()))
		return
	}

	_ = os.WriteFile(reportFile, data, 0644)
}
