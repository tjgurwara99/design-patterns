package factory

// DatabaseDriverSub is a sub for database driver.
type DatabaseDriverSub int

// ProductMySQLService is a concrete ProductCreator.
type ProductMySQLService struct {
	db *DatabaseDriverSub
}

// NewMySQLProductService returns a new ProductCreator. The reason why this is
// a factory method is that we have a concept of a "package" in Go. It C++, the
// pattern is used with a "Creator" class and a method would usually deal with
// the creation of another Object.
func NewMySQLProductService(db *DatabaseDriverSub) ProductCreator {
	return &ProductMySQLService{db}
}

var _ ProductCreator = (*ProductMySQLService)(nil)

// Create returns a new Product.
func (p *ProductMySQLService) Create(name string, price float64) *Product {
	return &Product{
		ID:    1,
		Name:  "MySQL Product",
		Price: 0.0,
	}
}
