package domain

type WishPriority string

const (
	PriorityHigh   WishPriority = "high"
	PriorityMedium WishPriority = "medium"
	PriorityLow    WishPriority = "low"
)

type Wish struct {
	ChildID  int          `json:"child_id"`
	ItemIDs  []int        `json:"item_ids"`
	Priority WishPriority `json:"priority"`
}

type WishRepository interface {
	GetByChildID(childID int) ([]Wish, error)
	GetAll() ([]Wish, error)
}
