package main

import (
	"fmt"
)

func main() {
	//create a tic toc toe board
	var s []int
	fmt.Println(s)

	s = append(s, 0)
	fmt.Println(s)

	s = append(s, 1)
	fmt.Println(s)

	s = append(s, 2, 3, 4, 5)
	fmt.Println(s)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Println("hello world!")
}
