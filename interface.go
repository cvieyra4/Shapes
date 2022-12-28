package shapes

import "fmt"

type geometric interface {
	area() float64
	perimeter() float64
}

func Measures(g geometric) {
	fmt.Println("Measures:", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perimeter())
}
