package main

import (
	"fmt"
)

func main() {
	var a int = 10
	fmt.Println("a :", &a)
	var p *int = &a
	fmt.Println("p1 :", p)
	fmt.Println("p2 :", &p)
	fmt.Println("p3 :", *p)
	fmt.Println("p4 :", *(&a))


}