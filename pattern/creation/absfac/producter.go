package absfac

// ProductCreator is the interface of product repository.
type ProductCreator interface {
	Create(name string, price float64) *Product
}
