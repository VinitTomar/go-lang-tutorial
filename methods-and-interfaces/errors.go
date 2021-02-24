package methodsandinterfaces

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x ErrNegativeSqrt) (float64, error) {

	if x < 0 {
		return 0, x
	}

	z := float64(x)
	v := float64(x)

	for i := 0; i < 1000; i++ {
		diff := z*z - v
		if diff == 0 {
			break
		}

		z -= diff / (2 * z)
	}

	return z, nil
}

func (f ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(f))
}

// ShowError ...
func ShowError() {
	println("======================== Custom errors =====================")

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
