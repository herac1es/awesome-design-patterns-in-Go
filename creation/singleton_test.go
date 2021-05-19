package creation

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 100; i++ {
		if GetInstanceLazy() != GetInstanceLazy() {
			t.Fatal("not pass")
		}
	}

	for i := 0; i < 100; i++ {
		if GetInstanceHungry() != GetInstanceHungry() {
			t.Fatal("not pass")
		}
	}

	for i := 0; i < 100; i++ {
		if GetInstanceLazyII() != GetInstanceLazyII() {
			t.Fatal("not pass")
		}
	}
}

func BenchmarkGetInstanceHungry(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceHungry() != GetInstanceHungry() {
				b.Errorf("test fail")
			}
		}
	})
}
func BenchmarkGetInstanceLazy(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazy() != GetInstanceLazy() {
				b.Errorf("test fail")
			}
		}
	})
}

func BenchmarkGetInstanceLazyII(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazyII() != GetInstanceLazyII() {
				b.Errorf("test fail")
			}
		}
	})
}
