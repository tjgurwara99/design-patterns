package absfac_test

import (
	"testing"

	"github.com/tjgurwara99/design-patterns/pattern/creational/absfac"
)

func testAddProduct(
	t *testing.T,
	af absfac.AbstractFactory,
	db *absfac.DatabaseDriverSub,
	expectedProduct *absfac.Product,
) {
	service := af(db)
	newProduct := service.Create("New Product", 0.0)
	if newProduct.Name != expectedProduct.Name {
		t.Errorf("Expected %v, got %v", expectedProduct, newProduct)
	}
}

func TestAddProductMongo(t *testing.T) {
	db := absfac.DatabaseDriverSub(1)
	expectedProduct := &absfac.Product{
		ID:    1,
		Name:  "Mongo Product",
		Price: 0.0,
	}
	testAddProduct(t, absfac.NewMongoProductService, &db, expectedProduct)
}

func TestAddProductMySQL(t *testing.T) {
	db := absfac.DatabaseDriverSub(2)
	expectedProduct := &absfac.Product{
		ID:    1,
		Name:  "MySQL Product",
		Price: 0.0,
	}
	testAddProduct(t, absfac.NewMySQLProductService, &db, expectedProduct)
}

func TestAddProductTestService(t *testing.T) {
	db := absfac.DatabaseDriverSub(3)
	expectedProduct := &absfac.Product{
		ID:    1,
		Name:  "Test Product",
		Price: 0.0,
	}
	testAddProduct(t, absfac.NewTestProductService, &db, expectedProduct)
}
