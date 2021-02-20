package moretypes

// ShowPointer ...
func ShowPointer() {
	println("======================== Pointers =====================")

	i, j := 42, 2701

	p := &i
	println(*p)
	*p = 21
	println(i)

	p = &j
	*p = *p / 37
	println(j)

}
