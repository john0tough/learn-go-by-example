package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Rectangle{
		Width:  10.0,
		Height: 10.0,
	}.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		Name    string
		Shape   Shape
		hasArea float64
	}{
		{
			Name: "Rectangle",
			Shape: Rectangle{
				Width:  12,
				Height: 6,
			},
			hasArea: 72.0,
		},
		{
			Name: "Circle",
			Shape: Circle{
				Radiuos: 10,
			},
			hasArea: 314.1592653589793,
		},
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	}

	for _, tt := range areaTest {
		t.Run(tt.Name, func(t *testing.T) {
			checkArea(t, tt.Shape, tt.hasArea)
		})
	}
}
