package flowcontrol

import "fmt"

func mySqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 1000; i++ {
		diff := z*z - x
		if diff == 0 {
			break
		}

		z -= diff / (2 * z)
	}

	return z
}

// ShowMySqrt ...
func ShowMySqrt() {
	println("======================== Custom square root =====================")

	for i := 10000; i < 10010; i++ {
		fmt.Printf("Sqrt of %v = %v\n", i, mySqrt(float64(i)))
	}

	println()
}
