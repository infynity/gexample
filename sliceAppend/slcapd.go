package main

import "fmt"

func main(){
	fmt.Println('被',string(34987))
	var sd []int
	fmt.Println(sd,sd==nil)
	fmt.Printf("%p\n",&sd)

	sd =make([]int,0)
	fmt.Println(sd,sd==nil)
	fmt.Printf("%p\n",&sd)

	//演示切片的使用 make
	var slice []float64 = make([]float64, 5, 500)
	fmt.Printf("%p\n",&slice[0])
	fmt.Println(len(slice),cap(slice))
	slice = append(slice,1.2)

	fmt.Printf("%p\n",&slice[0])
	fmt.Println(len(slice),cap(slice))

	slice2 := append(slice,1.23)
	fmt.Printf("%p\n",&slice2[0])
	fmt.Println(len(slice2),cap(slice2))
}