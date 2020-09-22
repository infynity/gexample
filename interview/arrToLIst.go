package main

import (
	"errors"
	"fmt"
)

func main() {

	rawMock :=mockList{front: -1,rear: -1,max: 5}


	fmt.Println(rawMock.arr)
}


type mockList struct {

	front int
	rear int

	arr [5]int
	max int
}
func  (this *mockList) pop()(int,error){

	if this.front==this.rear{
		return -2,errors.New("empty queue")
	}
	this.front = this.front+1

	if this.front<=this.rear{
		num:=this.arr[this.front]
		return num,nil
	}else{
		return -1,errors.New("nothing to pop ")

	}



}

func (this *mockList) push(num int) error{

	if this.rear==this.max-1{
		fmt.Println("queue is full")
		return errors.New("qwe")
	}

	this.rear++
	this.arr[this.rear]=num
	return nil
}


func (this *mockList)show(){

	for i:=this.front+1;i<=this.rear;i++{
		fmt.Println(this.arr[i])
	}

}