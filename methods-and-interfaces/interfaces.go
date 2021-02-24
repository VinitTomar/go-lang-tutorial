package methodsandinterfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Myfloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

// ShowInterface ...
func ShowInterface() {
	println("======================== Interfaces =====================")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v

	fmt.Println(a.Abs())

	println("======================== Empty Interfaces =====================")

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	var j interface{} = "hello"

	s := j.(string)
	fmt.Println(s)

	s, ok := j.(string)
	fmt.Println(s, ok)

	g, ok := j.(float32)
	fmt.Println(g, ok)

	println("======================== Type switch =====================")

	do(21)
	do("hello")
	do(true)

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
