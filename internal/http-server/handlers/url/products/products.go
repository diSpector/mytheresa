package products

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/diSpector/mytheresa.git/internal/cache"
	"github.com/diSpector/mytheresa.git/internal/discount"
	"github.com/diSpector/mytheresa.git/internal/domain"
	"github.com/diSpector/mytheresa.git/internal/storage"
	"github.com/diSpector/mytheresa.git/internal/validators"
)

const (
	CACHE_PREFIX          = `products`
	CACHE_PREFIX_ALL      = `all`
	CACHE_PREFIX_CATEGORY = `category`
	CACHE_PREFIX_PRICE    = `price`
)

func New(ctx context.Context, store storage.Storage, cache cache.Cache, discounts discount.Discounts) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var err error
		var products []domain.Product

		category := req.URL.Query().Get(`category`)
		priceLessThan := req.URL.Query().Get(`priceLessThan`)

		// if priceLessThan is invalid, it will be equal 0 (no filtering),
		// or it is also possible to return http error 400 (Bad Request) if neccessary;

		var priceLessThanInt = 0
		if validators.ValidatePositiveInt(priceLessThan) {
			// err could be omitted, because validation is successful
			priceLessThanInt, _ = strconv.Atoi(priceLessThan)
		}

		resp.Header().Set("Content-Type", "application/json")

		var ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		cacheKey := getCacheKey(category, priceLessThanInt)

		if productsFromCache, isFound, err := cache.Get(ctx, cacheKey); err != nil {
			// if err found in cache, just log it and go to the db routine
			log.Printf(`err get value for key="%s" from cache: %s`, cacheKey, err)
		} else if isFound {
			// if found in cache, return the value immediately
			err = json.Unmarshal([]byte(productsFromCache), &products)
			if err == nil {
				resp.WriteHeader(http.StatusOK)
				json.NewEncoder(resp).Encode(products)
				log.Printf(`got data for key=%s from cache`, cacheKey)
				return
			} else {
				log.Printf(`err unmarshal val for key=%s from cache: %s`, cacheKey, err)
			}
		}

		var productsDb []storage.Product

		// get products from db
		if priceLessThanInt == 0 && category == `` {
			productsDb, err = store.GetProducts(ctx)
		} else if priceLessThanInt > 0 && category == `` {
			productsDb, err = store.GetProductsUnderPrice(ctx, priceLessThanInt)
		} else if priceLessThanInt == 0 && category != `` {
			productsDb, err = store.GetProductsByCategory(ctx, category)
		} else {
			productsDb, err = store.GetProductsByCategoryUnderPrice(ctx, category, priceLessThanInt)
		}

		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			log.Printf(`err get products for category = %s, priceLessThan = %s: %s`, category, priceLessThan, err)
		}

		// for every found product convert to domain model and calculate discount
		for i := range productsDb {
			product := domain.NewProduct(productsDb[i].Sku, productsDb[i].Name, productsDb[i].Category, productsDb[i].Price)
			discount := discounts.CalcTotalDiscount(product)
			product.ApplyDiscount(discount)
			products = append(products, product)
		}

		// return OK code and list of products
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(products)

		// add found products to cache
		productsStr, err := json.Marshal(products)
		if err != nil {
			log.Println(`err marshal products:`, err)
		} else {
			err = cache.Set(ctx, cacheKey, string(productsStr))
			if err != nil {
				log.Printf(`err set cache for key = %s: %s`, cacheKey, err)
			}
		}
	}
}

func getCacheKey(category string, price int) string {
	if price == 0 && category == `` {
		// products:all
		return fmt.Sprintf("%s:%s", CACHE_PREFIX, CACHE_PREFIX_ALL)
	} else if price > 0 && category == `` {
		// products:price.val
		return fmt.Sprintf("%s:%s.%d", CACHE_PREFIX, CACHE_PREFIX_PRICE, price)
	} else if price == 0 && category != `` {
		// products:category.val
		return fmt.Sprintf("%s:%s.%s", CACHE_PREFIX, CACHE_PREFIX_CATEGORY, category)
	} else {
		// products:category.val:price.val
		return fmt.Sprintf("%s:%s.%s:%s.%d", CACHE_PREFIX, CACHE_PREFIX_CATEGORY, category, CACHE_PREFIX_PRICE, price)
	}
}
