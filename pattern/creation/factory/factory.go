package factory

type ProductCreator interface {
	Create(name string, price float64) *Product
}
