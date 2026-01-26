```go
package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"giftcalc/internal/infrastructure/logger"

	"github.com/spf13/cobra"
)

// Конфигурация приложения
type config struct {
	dataDir   string
	verbose   bool
	logLevel  string
	logFormat string // "json" или "console"
	logToFile string
}

var cfg = &config{}

var rootCmd = &cobra.Command{
	Use:   "giftcalc",
	Short: "Система расчета новогодних подарков",
	Long: `GiftCalc CLI - инструмент Аналитического отдела Деда Мороза
для расчета стоимости и комплектации индивидуальных подарков для детей.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupSignalHandling()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Синхронизация логгера при завершении, ошибки игнорируются
		_ = logger.Sync()
	},
}

// Execute запускает корневую команду и обрабатывает ошибки выполнения
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Ошибка выполнения команды",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
}

func init() {
	initFlags()
}

// initFlags инициализирует глобальные флаги приложения
func initFlags() {
	rootCmd.PersistentFlags().StringVarP(&cfg.dataDir, "data-dir", "d", "./data", 
		"Каталог с данными")
	rootCmd.PersistentFlags().BoolVarP(&cfg.verbose, "verbose", "v", false, 
		"Подробный вывод")
	rootCmd.PersistentFlags().StringVar(&cfg.logLevel, "log-level", "info",
		"Уровень логирования (debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().StringVar(&cfg.logFormat, "log-format", "console",
		"Формат логов (json, console)")
	rootCmd.PersistentFlags().StringVar(&cfg.logToFile, "log-file", "",
		"Файл для записи логов (если не указан - stdout)")
}

// setupSignalHandling настраивает обработку сигналов для graceful shutdown
func setupSignalHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go handleSignals(sigChan)
}

// handleSignals обрабатывает полученные сигналы завершения
func handleSignals(sigChan <-chan os.Signal) {
	sig := <-sigChan
	logger.Sugar().Infow("Получен сигнал завершения",
		"signal", sig.String(),
	)
	os.Exit(0)
}
```