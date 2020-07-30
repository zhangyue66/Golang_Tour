package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var yz map[string]string

func main() {

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{14.1111, 56.765}

	yz = make(map[string]string)

	yz["yue zhang"] = "handsome"

	fmt.Println(m["Bell Labs"])
	fmt.Println(yz["yue zhang"])

}
