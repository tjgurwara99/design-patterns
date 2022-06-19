package factory

import (
	"testing"
)

func TestNewProductor(t *testing.T) {
	db := DBSub(1)
	t.Run("Mysql test", func(t *testing.T) {
		productor := NewProductor(&db)
		_, ok := productor.(*mysqlProductor)
		if !ok {
			t.Errorf("expected type: *factory.mysqlProductor, got: %T", productor)
		}
	})
	t.Run("In Memory test", func(t *testing.T) {
		productor := NewProductor(nil)
		_, ok := productor.(*inMemoryProductor)
		if !ok {
			t.Errorf("expected type: *factory.inMemoryProductor, got: %T", productor)
		}
	})
}
