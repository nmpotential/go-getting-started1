package main

import "fmt"

func main()  {
	m := map[string]int{"foo":42}
	fmt.Println(m)

	//Print specific element of map using key type i.e. "foo"
	fmt.Println(m["foo"])

	//Reassign value within map
	m["foo"] = 27
	fmt.Println(m)

	//Delete value within map
	delete(m, "foo")
	fmt.Println(m)
}
