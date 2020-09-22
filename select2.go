package main

import (
	"fmt"
	"time"
)

func chantest1(obj chan int){
	for i:=0;i<=5;i++{
		obj<-i
	}
}


func chantest2(obj chan int){
	for i:=0;i<=5;i++{
		obj<-i
	}
}

func main(){

	chan1:=make(chan int ,5)
	chan2:=make(chan int ,5)

	go chantest1(chan1)
	go chantest2(chan2)

	for{
		select {
			case v1:=<-chan1:

				fmt.Println(v1)
			case v2:=<-chan2:

				fmt.Println(v2)
			default:
				fmt.Println(123)
				time.Sleep(1e9)
		}
	}
}
