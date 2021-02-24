package methodsandinterfaces

import (
	"fmt"
	"strconv"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
	return p.Name + " ( " + strconv.Itoa(p.Age) + " years"
}

// ShowStringer ...
func ShowStringer() {
	println("======================== Stringer =====================")

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
