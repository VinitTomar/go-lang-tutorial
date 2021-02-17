package main

import (
	"fmt"
	"math/rand"

	L "./code"
)

func main() {
	L.Say()
	fmt.Println("My favorite number is: ", rand.Intn(10))
	fmt.Println("Sum of 3 & 4 is: ", L.Add(3, 4))

	a, b := L.Swap(55, 11)

	fmt.Println("Swap ", a, b)
}
