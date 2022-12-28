package shapes

import "math"

type Circle struct {
	Radio float64
}

func (cir *Circle) area() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Circle) perimeter() float64 {
	return 2 * math.Pi * cir.Radio
}
