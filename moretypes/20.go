package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"bell labs": Vertex{
		12.11, 12.44,
	},
	"google": Vertex{
		55.777, 12.999,
	},
}

func main() {
	fmt.Println(m)
}
