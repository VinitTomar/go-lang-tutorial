package moretypes

import (
	"fmt"
	"strings"
)

type Vertex2 struct {
	lat, lng float64
}

var m map[string]Vertex2

var m2 = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
}

var m3 = map[string]Vertex2{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

// ShowMap ...
func ShowMap() {
	println("======================== Map =====================")

	m = make(map[string]Vertex2)

	m["lab1"] = Vertex2{
		40.68433, -74.39967,
	}

	fmt.Println(m["lab1"])
	fmt.Println(m2)
	fmt.Println(m3)

	s := "asdf asdf asdfd"

	sub := strings.Fields(s)

	mp := make(map[string]int)

	for i := range sub {
		str := sub[i]
		mp[str]++
	}

	fmt.Println(mp)

}
