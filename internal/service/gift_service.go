package service

import (
	"errors"
	"log/slog"
	"sync"
	"time"

	"giftcalc/internal/domain"
)

// Опции для GiftService
type GiftServiceOption func(*GiftService)

func WithMetrics(enabled bool) GiftServiceOption {
	return func(s *GiftService) {
		s.enableMetrics = enabled
	}
}

type GiftService struct {
	childRepo  domain.ChildRepository
	giftRepo   domain.GiftRepository
	wishRepo   domain.WishRepository
	regionRepo domain.RegionRepository

	maxGiftPrice  float64
	enableMetrics bool

	mu sync.RWMutex
}

func NewGiftService(
	childRepo domain.ChildRepository,
	giftRepo domain.GiftRepository,
	wishRepo domain.WishRepository,
	regionRepo domain.RegionRepository,
	opts ...GiftServiceOption,
) *GiftService {
	service := &GiftService{
		childRepo:  childRepo,
		giftRepo:   giftRepo,
		wishRepo:   wishRepo,
		regionRepo: regionRepo,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

func (s *GiftService) SetMaxGiftPrice(price float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxGiftPrice = price

	slog.Info("Установлен лимит стоимости подарка",
		slog.Float64("max_price", price),
	)
}

func (s *GiftService) CalculateForAll() (*domain.Report, error) {
	slog.Info("Начало расчета подарков для всех детей")

	children, err := s.childRepo.GetAll()
	if err != nil {
		slog.Error("Ошибка получения списка детей",
			slog.String("err", err.Error()),
		)
		return nil, err
	}

	slog.Debug("Данные о детях получены",
		slog.Int("total_children", len(children)),
	)

	report := &domain.Report{
		GeneratedAt: time.Now(),
		Version:     "1.0.0",
		Results:     make([]domain.ChildResult, 0, len(children)),
		Summary: domain.ReportSummary{
			TotalChildren: len(children),
		},
	}

	// Настройка параллельной обработки
	workerCount := 10 // Можно вынести в конфиг
	if len(children) < workerCount {
		workerCount = len(children)
	}

	var wg sync.WaitGroup
	jobs := make(chan domain.Child, len(children))
	results := make(chan domain.ChildResult, len(children))

	// Запуск воркеров
	slog.Debug("Запуск воркеров для параллельной обработки",
		slog.Int("worker_count", workerCount),
	)

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go s.worker(i, jobs, results, &wg)
	}

	// Отправка заданий
	for _, child := range children {
		jobs <- child
	}
	close(jobs)

	// Ожидание завершения
	go func() {
		wg.Wait()
		close(results)
	}()

	// Сбор результатов
	var (
		totalCost   float64
		minCost     = 1000000.0 // Большое начальное значение
		maxCost     = 0.0
		totalWeight float64
		processed   int
	)

	for result := range results {
		report.Results = append(report.Results, result)
		processed++

		// Прогресс каждые 1000 обработанных детей
		if processed%1000 == 0 {
			slog.Info("Прогресс расчета",
				slog.Int("processed", processed),
				slog.Int("total", len(children)),
				slog.Float64("progress_percent", float64(processed)/float64(len(children))*100),
			)
		}

		if result.Error == "" {
			report.Summary.Successful++
			totalCost += result.FinalCost
			totalWeight += result.TotalWeight

			// Обновление мин/макс стоимости
			if result.FinalCost < minCost {
				minCost = result.FinalCost
			}
			if result.FinalCost > maxCost {
				maxCost = result.FinalCost
			}
		} else {
			report.Summary.Failed++
			slog.Warn("Ошибка расчета подарка для ребенка",
				slog.Int("child_id", result.ChildID),
				slog.String("error", result.Error),
			)
		}
	}

	// Итоговая статистика
	report.TotalCost = totalCost
	report.Summary.MinCost = minCost
	report.Summary.MaxCost = maxCost
	report.Summary.TotalWeight = totalWeight
	if report.Summary.Successful > 0 {
		report.Summary.AvgCostPerGift = totalCost / float64(report.Summary.Successful)
	}

	slog.Info("Расчет завершен",
		slog.Int("total", report.Summary.TotalChildren),
		slog.Int("successful", report.Summary.Successful),
		slog.Int("failed", report.Summary.Failed),
		slog.Float64("total_cost", report.TotalCost),
		slog.Float64("avg_cost", report.Summary.AvgCostPerGift),
	)

	return report, nil
}

func (s *GiftService) worker(id int, jobs <-chan domain.Child, results chan<- domain.ChildResult, wg *sync.WaitGroup) {
	defer wg.Done()

	workerLogger := slog.With(slog.Int("worker_id", id))
	workerLogger.Debug("Воркер запущен")

	for child := range jobs {
		results <- s.calculateForChild(child)
	}

	workerLogger.Debug("Воркер завершил работу")
}

func (s *GiftService) calculateForChild(child domain.Child) domain.ChildResult {
	result := domain.ChildResult{
		ChildID:   child.ID,
		ChildName: child.Name,
		Region:    child.Region,
	}

	// Контекстный логгер для конкретного ребенка
	childLogger := slog.With(
		slog.Int("child_id", child.ID),
		slog.String("child_name", child.Name),
		slog.Int("age", child.Age),
		slog.String("region", child.Region),
	)

	childLogger.Debug("Начало расчета подарка")

	// Получение пожеланий
	var wishes []domain.Wish
	if s.wishRepo != nil {
		var err error
		wishes, err = s.wishRepo.GetByChildID(child.ID)
		if err != nil {
			childLogger.Warn("Ошибка получения пожеланий",
				slog.String("err", err.Error()),
			)
		} else {
			childLogger.Debug("Пожелания получены",
				slog.Int("wish_count", len(wishes)),
			)
		}
	}

	// Подбор подарков
	selections, err := s.selectGiftsForChild(child, wishes)
	if err != nil {
		result.Error = err.Error()
		childLogger.Error("Ошибка подбора подарков",
			slog.String("err", err.Error()),
		)
		return result
	}

	result.Selections = selections
	childLogger.Debug("Подарки подобраны",
		slog.Int("item_count", len(selections)),
	)

	// Расчет стоимости
	for _, selection := range selections {
		result.TotalCost += selection.Item.Price
		result.TotalWeight += selection.Item.Weight
	}

	// Применение регионального коэффициента
	coefficient := 1.0
	if s.regionRepo != nil {
		if coeff, err := s.regionRepo.GetCoefficient(child.Region); err == nil {
			coefficient = coeff
		} else {
			childLogger.Warn("Региональный коэффициент не найден, используется 1.0",
				slog.String("region", child.Region),
				slog.String("err", err.Error()),
			)
		}
	}

	result.Coefficient = coefficient
	result.FinalCost = result.TotalCost * coefficient

	childLogger.Debug("Стоимость рассчитана",
		slog.Float64("base_cost", result.TotalCost),
		slog.Float64("coefficient", coefficient),
		slog.Float64("final_cost", result.FinalCost),
		slog.Float64("total_weight", result.TotalWeight),
	)

	// Проверка максимальной стоимости
	if s.maxGiftPrice > 0 && result.FinalCost > s.maxGiftPrice {
		childLogger.Warn("Стоимость подарка превышает лимит",
			slog.Float64("cost", result.FinalCost),
			slog.Float64("limit", s.maxGiftPrice),
			slog.Float64("excess", result.FinalCost-s.maxGiftPrice),
		)

		// Попытка оптимизации
		if optimized, ok := s.optimizeGiftCost(selections, s.maxGiftPrice); ok {
			result.Selections = optimized
			// Пересчет стоимости после оптимизации
			result.TotalCost = 0
			result.TotalWeight = 0
			for _, selection := range optimized {
				result.TotalCost += selection.Item.Price
				result.TotalWeight += selection.Item.Weight
			}
			result.FinalCost = result.TotalCost * coefficient

			childLogger.Info("Подарок оптимизирован для соблюдения лимита",
				slog.Float64("new_cost", result.FinalCost),
			)
		}
	}

	childLogger.Debug("Расчет для ребенка завершен")

	return result
}

func (s *GiftService) selectGiftsForChild(
	child domain.Child,
	wishes []domain.Wish,
) ([]domain.GiftSelection, error) {
	var selections []domain.GiftSelection

	slog.Debug("Начало подбора подарков",
		slog.Int("wish_count", len(wishes)),
	)

	// Обработка пожеланий
	itemsSelected := make(map[int]bool) // Для избежания дублирования

	for _, wish := range wishes {
		for _, itemID := range wish.ItemIDs {
			if itemsSelected[itemID] {
				continue // Пропускаем дубли
			}

			item, err := s.giftRepo.FindByID(itemID)
			if err != nil {
				slog.Debug("Подарок из пожелания не найден в каталоге",
					slog.Int("item_id", itemID),
					slog.String("err", err.Error()),
				)
				continue
			}

			// Проверка возраста
			if child.Age < item.MinAge {
				slog.Debug("Подарок не подходит по возрасту",
					slog.Int("item_id", itemID),
					slog.Int("child_age", child.Age),
					slog.Int("item_min_age", item.MinAge),
				)
				continue
			}

			selection := domain.GiftSelection{
				Item:   *item,
				Reason: string(wish.Priority) + "_priority_wish",
			}
			selections = append(selections, selection)
			itemsSelected[itemID] = true

			slog.Debug("Подарок добавлен из пожелания",
				slog.Int("item_id", itemID),
				slog.String("item_name", item.Name),
				slog.String("reason", selection.Reason),
			)

			// Ограничиваем количество предметов
			if len(selections) >= 5 {
				slog.Debug("Достигнут лимит предметов в подарке (5)")
				return selections, nil
			}
		}
	}

	// Если пожеланий нет или их недостаточно, добавляем базовые подарки
	if len(selections) < 3 {
		slog.Debug("Недостаточно подарков из пожеланий, добавление базовых",
			slog.Int("current_count", len(selections)),
		)

		basicItems, err := s.giftRepo.FindByCategory("standard")
		if err != nil {
			return nil, err
		}

		for _, item := range basicItems {
			if child.Age >= item.MinAge && len(selections) < 5 && !itemsSelected[item.ID] {
				selection := domain.GiftSelection{
					Item:   item,
					Reason: "basic_gift",
				}
				selections = append(selections, selection)
				itemsSelected[item.ID] = true

				slog.Debug("Базовый подарок добавлен",
					slog.Int("item_id", item.ID),
					slog.String("item_name", item.Name),
				)

				if len(selections) >= 5 {
					break
				}
			}
		}
	}

	if len(selections) == 0 {
		slog.Error("Не удалось подобрать подходящие подарки")
		return nil, errors.New("no suitable gifts")
	}

	slog.Debug("Подбор подарков завершен",
		slog.Int("total_items", len(selections)),
	)

	return selections, nil
}

func (s *GiftService) optimizeGiftCost(
	selections []domain.GiftSelection,
	maxPrice float64,
) ([]domain.GiftSelection, bool) {
	// TODO: Реализовать логику оптимизации стоимости
	// Это может быть: замена дорогих предметов на более дешевые аналоги,
	// удаление наименее важных предметов и т.д.

	slog.Debug("Оптимизация стоимости не реализована")
	return selections, false
}
