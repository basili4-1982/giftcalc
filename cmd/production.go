package main

import "github.com/spf13/cobra"

var productionCmd = &cobra.Command{
	Use:   "production",
	Short: "Генерация производственного плана",
	Run:   runProduction,
}

func runProduction(cmd *cobra.Command, args []string) {

}
