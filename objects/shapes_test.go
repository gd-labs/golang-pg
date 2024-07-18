package objects

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("gpt %g, want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		checkArea(t, r, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})
}
