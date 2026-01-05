package domain

type Region struct {
	Name        string  `json:"name"`
	Coefficient float64 `json:"coefficient"`
}

type RegionRepository interface {
	GetCoefficient(regionName string) (float64, error)
	GetAll() ([]Region, error)
}
