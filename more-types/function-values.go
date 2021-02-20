package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// ShowFunctionValues ...
func ShowFunctionValues() {
	println("======================== Function values =====================")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println("Hypot: ", hypot(5, 12))
	fmt.Println("Compute with Hypot: ", compute(hypot))
	fmt.Println("Compute with math.Pow: ", compute(math.Pow))

	println("======================== Function closure =====================")

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(i, pos(i), neg(-2*i))
	}

}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}
