package main

import (
	"encoding/json"
	"giftcalc/internal/domain"
	"log/slog"
	"os"

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

	data, err := os.ReadFile(childrenDataFile)
	if err != nil {
		slog.Error("Не могу найти файл '" + childrenDataFile + "'")
		return
	}

	res := domain.ChildrenData{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		slog.Error("Не могу разобрать файл '"+childrenDataFile+"'", slog.String("err", err.Error()))
		return
	}

	for _, child := range res.Children {
		_ = child
	}
}
