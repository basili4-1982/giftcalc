```go
package domain

// WishPriority определяет приоритет желания
type WishPriority string

const (
	PriorityHigh   WishPriority = "high"
	PriorityMedium WishPriority = "medium"
	PriorityLow    WishPriority = "low"
)

// IsValid проверяет, является ли приоритет допустимым значением
func (p WishPriority) IsValid() bool {
	switch p {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return true
	default:
		return false
	}
}

// Wish представляет желание ребенка
type Wish struct {
	ChildID  int          `json:"child_id"`
	ItemIDs  []int        `json:"item_ids"`
	Priority WishPriority `json:"priority"`
}

// Validate проверяет корректность данных желания
func (w *Wish) Validate() error {
	if w.ChildID <= 0 {
		return ErrInvalidChildID
	}
	
	if len(w.ItemIDs) == 0 {
		return ErrEmptyItemIDs
	}
	
	if !w.Priority.IsValid() {
		return ErrInvalidPriority
	}
	
	return nil
}

// WishRepository определяет контракт для работы с хранилищем желаний
type WishRepository interface {
	// GetByChildID возвращает все желания для указанного ребенка
	GetByChildID(childID int) ([]Wish, error)
	
	// GetAll возвращает все желания из хранилища
	GetAll() ([]Wish, error)
}

// Ошибки доменной логики
var (
	ErrInvalidChildID  = errors.New("invalid child ID")
	ErrEmptyItemIDs    = errors.New("item IDs cannot be empty")
	ErrInvalidPriority = errors.New("invalid priority value")
)

// Инициализация пакета
func init() {
	// Проверка, что все константы приоритетов являются валидными
	priorities := []WishPriority{PriorityHigh, PriorityMedium, PriorityLow}
	for _, p := range priorities {
		if !p.IsValid() {
			panic(fmt.Sprintf("invalid priority constant: %s", p))
		}
	}
}
```