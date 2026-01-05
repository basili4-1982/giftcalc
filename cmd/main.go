package main

import (
	"fmt"
	"os"
)

func main() {
	// Регистрация команд
	rootCmd.AddCommand(
		calculateCmd,
		costCmd,
		productionCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
