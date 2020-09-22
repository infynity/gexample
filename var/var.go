package main

import "fmt"

func main(){
//a := make(chan int) creates unbuffered channel. That channel with zero buffer. You can send data through it.
	//
	//var a chan int creates channel variable and sets it to default value which is nil. And a nil channel is always blocking, that's why your program is deadlocked. You can not send data in nil channel.
	//
	//If you print the values, you will see the difference.



	jkl:=new(int)

	suda := new([]int)

	fmt.Println(jkl,suda,7777,*jkl,jkl==nil,*suda==nil,*suda)
	fmt.Printf("%p\n",suda)
	fmt.Println(&suda)
	return
	var ji int
	var sr []int
	fmt.Println(ji,sr,sr==nil)


		var i chan int
		fmt.Println(i)
		a := make(chan int)
		fmt.Println(a)


	var c []int

		fmt.Println(c)
		d := []int{}
	fmt.Println(d)
	var e =make([]int,0)
	fmt.Println(e)



}



