package domain

import "fmt"

const CUR_DEFAULT = `EUR`

type Product struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    Price  `json:"price"`
}

type Price struct {
	Original           int     `json:"original"`
	Final              int     `json:"final"`
	DiscountPercentage *string `json:"discount_percentage"`
	Currency           string  `json:"currency"`
}

func NewProduct(sku, name, category string, price int) Product {
	return Product{
		Sku:      sku,
		Name:     name,
		Category: category,
		Price: Price{
			Original: price,
			Final:    price,
			Currency: CUR_DEFAULT,
		},
	}
}

func (s *Product) ApplyDiscount(discount int) {
	if discount > 0 {
		// there is a possibility of int overflow, so the final price calculation method should depends on max possible price
		s.Price.Final = (100 - discount) * s.Price.Original / 100
		s.Price.DiscountPercentage = new(string)
		*s.Price.DiscountPercentage = fmt.Sprintf("%d%%", discount)
	}
}
