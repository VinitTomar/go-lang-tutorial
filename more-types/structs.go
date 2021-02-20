package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{2, 3}
	v2 = Vertex{X: 5}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

// ShowStruct ...
func ShowStruct() {
	println("======================== Struct =====================")
	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v)

	p := &v
	// (*p).X = 3e9         OR
	p.X = 2e4

	fmt.Println(v)
	fmt.Println(v1, v2, v3, p)

}
