package main

import (
	"fmt"
)

func getFibna(n int ) int {
	if n==1 ||  n==2 {
		return 1
	}

	return getFibna(n-1)+getFibna(n-2)
}


func monkey(n int) int{
	if n==1{
		return 1
	}
	//
	//if n==2{
	//	1+1  * 2 =4
	//}
	//
	//if n==3{
	//	4+1  *2 =10
	//}

	return (monkey(n-1)+1) *2

}

func main(){


	var as interface{}
	as ="asd"
	i,pk:=as.(string)
	v,ok:=(as.(int))
	fmt.Println(i,pk,v,ok)
	return
	fmt.Println(monkey(10))
	fmt.Println(getFibna(5))
}
