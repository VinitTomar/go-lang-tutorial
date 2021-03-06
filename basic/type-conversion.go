package basic

import "fmt"

// ShowTypeConversion ...
func ShowTypeConversion() {
	println("====================== Type Conversion =================")
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", u, u)
}
