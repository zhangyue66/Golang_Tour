package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, v := range pow {
		fmt.Println(v)
	}
}
