package absfac

// DatabaseDriverSub is a sub for database driver.
type DatabaseDriverSub int

type ProductMongoService struct {
	db *DatabaseDriverSub
}

var _ ProductCreator = (*ProductMongoService)(nil)

func NewMongoProductService(db *DatabaseDriverSub) ProductCreator {
	return &ProductMongoService{db}
}

func (p *ProductMongoService) Create(name string, price float64) *Product {
	return &Product{
		ID:    1,
		Name:  "Mongo Product",
		Price: 0.0,
	}
}

type ProductMySQLService struct {
	db *DatabaseDriverSub
}

var _ ProductCreator = (*ProductMySQLService)(nil)

func NewMySQLProductService(db *DatabaseDriverSub) ProductCreator {
	return &ProductMySQLService{db}
}

func (p *ProductMySQLService) Create(name string, price float64) *Product {
	return &Product{
		ID:    1,
		Name:  "MySQL Product",
		Price: 0.0,
	}
}

// ProductTestService is just a test service to substitute the actual database service.
type ProductTestService struct {
	db *DatabaseDriverSub
}

var _ ProductCreator = (*ProductTestService)(nil)

func NewTestProductService(db *DatabaseDriverSub) ProductCreator {
	return &ProductTestService{db}
}

func (p *ProductTestService) Create(name string, price float64) *Product {
	return &Product{
		ID:    1,
		Name:  "Test Product",
		Price: 0.0,
	}
}
