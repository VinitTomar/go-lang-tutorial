package flowcontrol

import (
	"fmt"
	"math"
)

// ShowIf ...
func ShowIf() {
	println("======================== If control statement =====================")

	fmt.Println("Sqrt of 4 is ", sqrt(4))
	fmt.Println("Sqrt of -4 is ", sqrt(-4))

}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}
