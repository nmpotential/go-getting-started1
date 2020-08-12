package main

import "fmt"

const (
	first = iota
	second = iota
)
const (
	third = iota
	fourth
)

func main()  {
	fmt.Println(first, second, third, fourth)
}