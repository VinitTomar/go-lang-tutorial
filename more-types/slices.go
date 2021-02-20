package moretypes

import (
	"fmt"
)

// ShowSlice ...
func ShowSlice() {
	println("======================== Slice =====================")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	println("======================== Slice literal with struct =====================")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	s1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s1)

	// Slice the slice to give it zero length.
	s1 = s1[:0]
	printSlice(s1)

	// Extend its length.
	s1 = s1[:4]
	printSlice(s1)

	// Drop its first two values.
	s1 = s1[2:]
	printSlice(s1)

	s1 = s1[:4]
	printSlice(s1)

	println("======================== Nil slice =====================")

	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}

	println("======================== Make slice =====================")

	a1 := make([]int, 5)
	printSlice2("a", a1)

	b1 := make([]int, 0, 5)
	printSlice2("b", b1)

	c1 := b1[:2]
	printSlice2("c", c1)

	d1 := c1[2:5]
	printSlice2("d", d1)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
