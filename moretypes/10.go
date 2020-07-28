package main

import "fmt"

func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	x := s[1:4]
	fmt.Println(x)

	p := s[:2]
	fmt.Println(p)

	q := s[2:]
	fmt.Println(q)

}
