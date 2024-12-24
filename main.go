package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/diSpector/mytheresa.git/internal/cache/rediscache"
	"github.com/diSpector/mytheresa.git/internal/config"
	"github.com/diSpector/mytheresa.git/internal/discount"
	"github.com/diSpector/mytheresa.git/internal/http-server/handlers/url/products"
	"github.com/diSpector/mytheresa.git/internal/storage/postgres"
)

const (
	DEFAULT_CONFIG_PATH = `./config/conf.yaml`
	SKU_DISCOUNT        = 15
	CATEGORY_DISCOUNT   = 30
)

func main() {
	// read config from -config flag or default value
	confPath := flag.String("config", DEFAULT_CONFIG_PATH, "set config path")
	flag.Parse()

	if confPath == nil {
		log.Fatal("err read config flag")
	}

	conf, err := config.Read(*confPath)
	if err != nil {
		log.Fatalf("err process config file: %s", err)
	}

	ctx := context.Background()

	// init storage: postgres
	storage, err := postgres.New(conf.Storage.Host, conf.Storage.Port, conf.Storage.User, conf.Storage.Password, conf.Storage.Database)
	if err != nil {
		log.Fatalf("err connect to storage: %s", err)
	}

	// init cache: redis
	cache := rediscache.New(conf.Cache.Host, conf.Cache.Port, conf.Cache.Password, conf.Cache.Database, conf.Cache.Ttl)

	// init discounts (we can add here all the discounts one by one)
	discounts := discount.NewDiscounts(
		discount.NewSkuDiscount([]string{`000003`}, SKU_DISCOUNT),
		discount.NewCategoryDiscount([]string{`boots`}, CATEGORY_DISCOUNT))

	// init router
	router := http.NewServeMux()
	router.HandleFunc(`/products`, products.New(ctx, storage, cache, discounts))

	// prepare and run server
	srv := &http.Server{
		Addr:         conf.HttpServer.Address,
		Handler:      router,
		ReadTimeout:  conf.HttpServer.Timeout,
		WriteTimeout: conf.HttpServer.Timeout,
		IdleTimeout:  conf.HttpServer.IdleTimeout,
	}

	log.Println("run server on:", conf.HttpServer.Address)
	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start server:", err)
	}
}
