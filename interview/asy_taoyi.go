package main

import "fmt"

var cc int
//go build -gcflags=-m taoyi.go
func test2() {
	a:= []int{1,2,3,4}
	a[1]=4

}
var tslice = make([]int,10)

var ko = make(chan int)
func test3() []int{
	a:= []int{1,2,3,4}
	a[1]=4


	b:=make(chan int)
	fmt.Println(b)

	fmt.Println(ko)
	return a
}

func test4() map[int]int{
	a:= make(map[int]int,2)
	//a = 1114

	return a
}
//var a string
func test5() *string{
	a:= "asdsad"
	//a = 1114

	return &a
}

func test7() []int{
	tslice[1]=123
	return tslice
}
func test6() string{
	a:= "asdsad"
	//a = 1114

	return a
}
func main(){
	//test2()
	test3()
	test4()

	test7()
}