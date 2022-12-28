package shapes

type Square struct {
	Side float64
}

func (cua *Square) area() float64 {
	return cua.Side * cua.Side
}

func (cua *Square) perimeter() float64 {
	return 4 * cua.Side
}
