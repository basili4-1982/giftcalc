package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"giftcalc/internal/infrastructure/logger"

	"github.com/spf13/cobra"
)

var (
	dataDir   string
	verbose   bool
	logLevel  string
	logFormat string // "json" или "console"
	logToFile string
)

var rootCmd = &cobra.Command{
	Use:   "giftcalc",
	Short: "Система расчета новогодних подарков",
	Long: `GiftCalc CLI - инструмент Аналитического отдела Деда Мороза
для расчета стоимости и комплектации индивидуальных подарков для детей.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Обработка сигналов для graceful shutdown
		setupSignalHandling()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Синхронизация логгера при завершении
		if err := logger.Sync(); err != nil {
			// Игнорируем ошибку при завершении
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Ошибка выполнения команды",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}
}

func init() {
	// Глобальные флаги
	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "d", "./data", "Каталог с данными")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Подробный вывод")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info",
		"Уровень логирования (debug, info, warn, error, fatal)")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "console",
		"Формат логов (json, console)")
	rootCmd.PersistentFlags().StringVar(&logToFile, "log-file", "",
		"Файл для записи логов (если не указан - stdout)")
}

func setupSignalHandling() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		logger.Sugar().Infow("Получен сигнал завершения",
			"signal", sig.String(),
		)

		os.Exit(0)
	}()
}
