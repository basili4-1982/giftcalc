package main

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var costCmd = &cobra.Command{
	Use:   "cost",
	Short: "Расчет стоимости подарков",
	Run:   runCost,
}

func runCost(cmd *cobra.Command, args []string) {
	slog.Info("Расчет стоимости")
}
