package prototype

type Cube struct {
	width, height, depth int
}

func NewCube(width, height, depth int) *Cube {
	return &Cube{width, height, depth}
}

func (c *Cube) Clone() Cuber {
	return &Cube{c.width, c.height, c.depth}
}

func (c *Cube) Area() int {
	return c.width * c.height * c.depth
}
