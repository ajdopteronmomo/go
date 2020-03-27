package main

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	//a()
	c := a()
	c()
	c()
	// c()
	// a()
}
