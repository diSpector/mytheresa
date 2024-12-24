package discount

import (
	"fmt"
	"testing"

	"github.com/diSpector/mytheresa.git/internal/domain"
)

func TestSkuDiscount(t *testing.T) {
	testTable := []struct {
		Name     string
		Skus     []string
		Size     int
		Product  domain.Product
		Discount int
	}{
		{
			Name: `discount for 1 sku`,
			Skus: []string{`000003`},
			Size: 10,
			Product: domain.Product{
				Sku: `000003`,
			},
			Discount: 10,
		},
		{
			Name: `discount for 2 skus`,
			Skus: []string{`000003`, `000004`},
			Size: 10,
			Product: domain.Product{
				Sku: `000004`,
			},
			Discount: 10,
		},
		{
			Name: `without discount`,
			Skus: []string{`000003`, `000004`},
			Size: 10,
			Product: domain.Product{
				Sku: `000001`,
			},
			Discount: 0,
		},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("%s_%d", t.Name(), i), func(t *testing.T) {
			t.Log("testing:", test.Name)
			skuDiscount := NewSkuDiscount(test.Skus, test.Size)
			res := skuDiscount.CalcDiscount(test.Product)
			if res != test.Discount {
				t.Logf("should be - %v, got - %v", test.Discount, res)
				t.FailNow()
			}
		})
	}
}

func TestCategoryDiscount(t *testing.T) {
	testTable := []struct {
		Name       string
		Categories []string
		Size       int
		Product    domain.Product
		Discount   int
	}{
		{
			Name:       `discount for 1 category`,
			Categories: []string{`boots`},
			Size:       20,
			Product: domain.Product{
				Category: `boots`,
			},
			Discount: 20,
		},
		{
			Name:       `discount for 2 categories`,
			Categories: []string{`boots`, `sneakers`},
			Size:       10,
			Product: domain.Product{
				Category: `sneakers`,
			},
			Discount: 10,
		},
		{
			Name:       `without discount`,
			Categories: []string{`boots`, `sneakers`},
			Size:       10,
			Product: domain.Product{
				Category: `sweaters`,
			},
			Discount: 0,
		},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("%s_%d", t.Name(), i), func(t *testing.T) {
			t.Log("testing:", test.Name)
			categoryDiscount := NewCategoryDiscount(test.Categories, test.Size)
			res := categoryDiscount.CalcDiscount(test.Product)
			if res != test.Discount {
				t.Logf("should be - %v, got - %v", test.Discount, res)
				t.FailNow()
			}
		})
	}
}

func TestMultipleDiscounts(t *testing.T) {
	testTable := []struct {
		Name      string
		Discounts Discounts
		Product   domain.Product
		Discount  int
	}{
		{
			Name: `1 discount for category applied`,
			Discounts: NewDiscounts(
				NewCategoryDiscount([]string{`boots`}, 30),
				NewSkuDiscount([]string{`000003`}, 15),
			),
			Product: domain.Product{
				Category: `boots`,
				Sku:      `000001`,
			},
			Discount: 30,
		},
		{
			Name: `1 discount for sku applied`,
			Discounts: NewDiscounts(
				NewCategoryDiscount([]string{`boots`}, 30),
				NewSkuDiscount([]string{`000003`}, 15),
			),
			Product: domain.Product{
				Category: `sweaters`,
				Sku:      `000003`,
			},
			Discount: 15,
		},
		{
			Name: `discounts collide (max discount applied)`,
			Discounts: NewDiscounts(
				NewCategoryDiscount([]string{`boots`}, 30),
				NewSkuDiscount([]string{`000003`}, 15),
			),
			Product: domain.Product{
				Category: `boots`,
				Sku:      `000003`,
			},
			Discount: 30,
		},
		{
			Name: `without discounts`,
			Discounts: NewDiscounts(
				NewCategoryDiscount([]string{`boots`}, 30),
				NewSkuDiscount([]string{`000003`}, 15),
			),
			Product: domain.Product{
				Category: `sweaters`,
				Sku:      `000001`,
			},
			Discount: 0,
		},
	}

	for i, test := range testTable {
		t.Run(fmt.Sprintf("%s_%d", t.Name(), i), func(t *testing.T) {
			t.Log("testing:", test.Name)
			res := test.Discounts.CalcTotalDiscount(test.Product)
			if res != test.Discount {
				t.Logf("should be - %v, got - %v", test.Discount, res)
				t.FailNow()
			}
		})
	}
}
