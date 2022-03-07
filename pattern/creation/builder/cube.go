package builder

type Cube struct {
	length int
	width  int
	height int
}

func NewCube() *Cube {
	return &Cube{}
}

func (c *Cube) WithLength(length int) *Cube {
	c.length = length
	return c
}

func (c *Cube) WithWidth(width int) *Cube {
	c.width = width
	return c
}

func (c *Cube) WithHeight(height int) *Cube {
	c.height = height
	return c
}

func (c *Cube) Area() int {
	return c.length * c.width * c.height
}
