package absfac

func AddProduct(af AbstractFactory, db *DatabaseDriverSub) *Product {
	service := af(db)
	newProduct := service.Create("New Product", 0.0)
	return newProduct
}
