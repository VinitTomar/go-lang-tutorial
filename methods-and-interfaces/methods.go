package methodsandinterfaces

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (num MyFloat) abs2() float64 {
	if num < 0 {
		return float64(-num)
	}

	return float64(num)

}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ShowMethods ...
func ShowMethods() {
	println("======================== Methods =====================")

	v := Vertex{3, 4}

	v.scale(0.5)

	fmt.Println(v.abs())

	nm := MyFloat(-55)

	fmt.Println(nm.abs2())
}
