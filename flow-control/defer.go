package flowcontrol

// ShowDefer ...
func ShowDefer() {
	println("======================== Defer =====================")

	defer println("world!")

	defer print("Hello ")

	print("Okay ")

}
