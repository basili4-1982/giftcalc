```go
package main

import (
	"fmt"
	"os"
)

func main() {
	registerCommands()
	executeRootCommand()
}

// registerCommands регистрирует все дочерние команды в корневой команде
func registerCommands() {
	rootCmd.AddCommand(
		calculateCmd,
		costCmd,
		productionCmd,
	)
}

// executeRootCommand выполняет корневую команду и обрабатывает возможные ошибки
func executeRootCommand() {
	if err := rootCmd.Execute(); err != nil {
		handleExecutionError(err)
	}
}

// handleExecutionError обрабатывает ошибки выполнения команды
func handleExecutionError(err error) {
	fmt.Fprintln(os.Stderr, "Ошибка выполнения:", err)
	os.Exit(1)
}
```