package prototype

import "testing"

func TestArea(t *testing.T) {
	c := NewCube(10, 20, 30)
	if c.Area() != 6000 {
		t.Errorf("Area of cube is %d, want 6000", c.Area())
	}
}

func TestCubeCloner(t *testing.T) {
	c := NewCube(10, 20, 30)
	c2 := c.Clone()
	if c2.Area() != c.Area() {
		t.Errorf("Area of cube is %d, want %d", c2.Area(), c.Area())
	}
	if c == c2 {
		t.Errorf("c and c2 are the same instance and not clones")
	}
}
