package discount

import (
	"slices"

	"github.com/diSpector/mytheresa.git/internal/domain"
)

type Discounts []DiscountType

// NewDiscounts constructor for different types of discounts
func NewDiscounts(discounts ...DiscountType) Discounts {
	return Discounts(discounts)
}

func (s Discounts) CalcTotalDiscount(product domain.Product) int {
	var totalDiscount int

	// apply each type of discount to product and return maximum
	for i := range s {
		discount := s[i].CalcDiscount(product)
		if discount > totalDiscount {
			totalDiscount = discount
		}
	}

	return totalDiscount
}

type DiscountType interface {
	CalcDiscount(domain.Product) int
}

// discount for sku
type SkuDiscount struct {
	skus []string
	size int
}

func NewSkuDiscount(skus []string, size int) SkuDiscount {
	return SkuDiscount{
		skus: skus,
		size: size,
	}
}

func (s SkuDiscount) CalcDiscount(product domain.Product) int {
	if slices.Contains(s.skus, product.Sku) {
		return s.size
	}

	return 0
}

// discount for categories
type CategoryDiscount struct {
	categories []string
	size       int
}

func NewCategoryDiscount(categories []string, size int) CategoryDiscount {
	return CategoryDiscount{
		categories: categories,
		size:       size,
	}
}

func (s CategoryDiscount) CalcDiscount(product domain.Product) int {
	if slices.Contains(s.categories, product.Category) {
		return s.size
	}

	return 0
}
