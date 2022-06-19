// Package factory represents the use of the Factory Method pattern.
// This is a very general pattern, I've been using it for years without
// knowing that its actually a well established pattern.
package factory

import (
	"errors"
	"fmt"
)

// Product is an example product type - your business entity (in onion architecture.)
type Product struct {
	Name  string
	Price int64 // the correct way to handle prices - floats are error prone
}

// Productor is a general interface for general db queries.
type Productor interface {
	CreateProduct(p *Product) error
	GetProduct(name string) (*Product, error)
}

func NewProductor(db *DBSub) Productor {
	if db == nil {
		return &inMemoryProductor{}
	}
	return &mysqlProductor{db: db}
}

type DBSub int

type mysqlProductor struct {
	db       *DBSub
	products []Product
}

func (m *mysqlProductor) CreateProduct(p *Product) error {
	m.products = append(m.products, *p)
	return nil
}

var ErrProductNotFound error = errors.New("DB product not found")

func (m *mysqlProductor) GetProduct(name string) (*Product, error) {
	for _, prod := range m.products {
		if prod.Name == name {
			return &Product{
				Name:  prod.Name,
				Price: prod.Price,
			}, nil
		}
	}
	return nil, fmt.Errorf("product not found: %w", ErrProductNotFound)
}

type inMemoryProductor struct {
	products []Product
}

func (i *inMemoryProductor) CreateProduct(p *Product) error {
	i.products = append(i.products, *p)
	return nil
}

func (i *inMemoryProductor) GetProduct(name string) (*Product, error) {
	for _, prod := range i.products {
		if prod.Name == name {
			return &Product{
				Name:  prod.Name,
				Price: prod.Price,
			}, nil
		}
	}
	return nil, fmt.Errorf("product not found: %w", ErrProductNotFound)
}
