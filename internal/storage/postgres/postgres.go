package postgres

import (
	"context"
	"fmt"

	"github.com/diSpector/mytheresa.git/internal/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(host string, port int, user, password, dbName string) (Storage, error) {
	db, err := sqlx.Connect(`postgres`, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName))
	if err != nil {
		return Storage{}, err
	}

	return Storage{db: db}, nil
}

func (s Storage) GetProducts(ctx context.Context) ([]storage.Product, error) {
	sql := `select * from products limit 5`
	var res []storage.Product

	err := s.db.SelectContext(ctx, &res, sql)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s Storage) GetProductsByCategory(ctx context.Context, category string) ([]storage.Product, error) {
	sql := `select * from products where category = :category limit 5`

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"category": category,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []storage.Product
	for rows.Next() {
		var pr storage.Product
		err = rows.StructScan(&pr)
		if err != nil {
			return nil, err
		}
		res = append(res, pr)
	}

	return res, nil
}

func (s Storage) GetProductsUnderPrice(ctx context.Context, priceUnder int) ([]storage.Product, error) {
	sql := `select * from products where price <= :price limit 5`

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"price": priceUnder,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []storage.Product
	for rows.Next() {
		var pr storage.Product
		err = rows.StructScan(&pr)
		if err != nil {
			return nil, err
		}
		res = append(res, pr)
	}

	return res, nil
}

func (s Storage) GetProductsByCategoryUnderPrice(ctx context.Context, category string, priceUnder int) ([]storage.Product, error) {
	sql := `select * from products where category = :category and price < :price limit 5`

	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"category": category,
		"price":    priceUnder,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []storage.Product
	for rows.Next() {
		var pr storage.Product
		err = rows.StructScan(&pr)
		if err != nil {
			return nil, err
		}
		res = append(res, pr)
	}

	return res, nil
}
