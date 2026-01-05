package service

import (
	"math"

	"giftcalc/internal/domain"
)

type PriceCalculator struct {
	regionRepo domain.RegionRepository
}

func NewPriceCalculator(regionRepo domain.RegionRepository) *PriceCalculator {
	return &PriceCalculator{
		regionRepo: regionRepo,
	}
}

func (c *PriceCalculator) CalculateWithRegion(basePrice float64, region string) (float64, float64, error) {
	coefficient := 1.0

	if c.regionRepo != nil {
		coeff, err := c.regionRepo.GetCoefficient(region)
		if err == nil {
			coefficient = coeff
		}
	}

	finalPrice := math.Round(basePrice*coefficient*100) / 100
	return finalPrice, coefficient, nil
}

func (c *PriceCalculator) ApplyBudgetLimit(price, maxPrice float64) (float64, bool) {
	if maxPrice <= 0 {
		return price, false
	}

	if price <= maxPrice {
		return price, false
	}

	// TODO: Более сложная логика оптимизации
	return maxPrice, true
}
