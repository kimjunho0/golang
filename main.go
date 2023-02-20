package main

import "fmt"

type Person struct {
	Age int
}

func main() {
	fmt.Println("Hello Go World")
	p := Person{Age: 4}
	fmt.Println(p.Age)
	return
}
