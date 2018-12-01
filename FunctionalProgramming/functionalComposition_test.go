package FunctionalProgramming

import "testing"

func TestCompose(t *testing.T) {
	f := func(i Any) Any {
		return i.(int) + i.(int)
	}
	g := func(j Any) Any {
		return j.(int) * j.(int)
	}

	composed := compose(f, g)
	if composed(5) != 50 {
		t.Error("Result should be 50!")
	}

	composed = compose(g, f)
	if composed(5) != 100 {
		t.Error("Result should be 100!")
	}
}
