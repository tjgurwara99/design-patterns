package factory_test

import (
	"testing"

	"github.com/tjgurwara99/design-patterns/pattern/creation/factory"
)

func TestProductMySQLService(t *testing.T) {
	db := factory.DatabaseDriverSub(1)
	product := factory.NewMySQLProductService(&db).Create("MySQL Product", 0.0)
	if product.Name != "MySQL Product" {
		t.Errorf("Expected MySQL Product, got %s", product.Name)
	}
}
