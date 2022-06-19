package builder

type Builder interface {
	WithLength(length int) Builder
	WithWidth(width int) Builder
	WithHeight(height int) Builder
	Area() int
}
