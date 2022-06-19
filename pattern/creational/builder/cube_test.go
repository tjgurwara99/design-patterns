package builder_test

import (
	"testing"

	"github.com/tjgurwara99/design-patterns/pattern/creational/builder"
)

func TestCubeBuilder(t *testing.T) {
	cube := builder.NewCube().WithLength(10).WithWidth(20).WithHeight(30).Area()
	if cube != 6000 {
		t.Errorf("Expected area to be 6000, got %d", cube)
	}
}
