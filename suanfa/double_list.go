package main

import "fmt"

type Node struct{

	no int
	name string
	prev,next *Node
}


func InsertNode(head *Node,Nownode *Node){
	
	//var temp Node
	
	temp:= head
	for {
		if temp.next==nil{
			break
		}

		temp = temp.next
	}
	
	
	temp.next=Nownode
	
	Nownode.prev=temp

}

func DelNode(){

}

func ShowAll(temp *Node){

	fmt.Println(temp.no,"init")
	//temp := head
	//fmt.Println(144,"-----",temp,&temp)
	fmt.Println(144,"-----",&temp,&temp.next)


	if temp.next == nil{
		fmt.Println("nothing")
		return
	}

	for{
		fmt.Println(temp.next.no,temp.next.name,&temp.next)
		temp = temp.next

		if temp.next==nil{
			fmt.Println("done")
			break
		}


	}

	fmt.Println(4,"-----",temp,&temp)
}

func ShowAllReverse(head *Node){
	fmt.Println(head.no,"init")
	fmt.Println(244,"-----",&head,&head.next)

	if head.next==nil{
		fmt.Println("none list")
	}

	for {
		if head.next==nil{
			break
		}
		head = head.next

	}


	for {
		fmt.Println(head.no,head.name)


		head = head.prev
		if head.prev==nil{
			fmt.Println("done")
			break
		}
	}

fmt.Println(head)

}


func main(){

	//1. 先创建一个头结点,
	head := &Node{}
	//2. 创建一个新的
	hero1 := &Node{
		no:   1,
		name: "宋江",
	}
	hero2 := &Node{no: 2,
		name:     "卢俊义",
		}
	hero3 := &Node{

		no:   3,
		name: "林冲",
	}
	hero4 := &Node{

		no:   4,
		name: "infy",
	}
	hero5 := &Node{

		no:   5,
		name: "infy3",
	}
	InsertNode(head, hero1)
	InsertNode(head, hero2)
	InsertNode(head, hero3)
	InsertNode(head, hero4)
	InsertNode(head, hero5)
	fmt.Println(1,"-----",&head,&head.next)
	ShowAll(head)

	fmt.Println(2,"-------",&head,head.next)


	fmt.Println("===")

	ShowAllReverse(head)
	fmt.Println(3,"-------",&head)

	//sj:=[]int{1,2,3}
	//ptrval(&sj)
	//ptrval(&sj)
	//
	//fmt.Println(sj)
}


func ptrval(j *[]int){

	fmt.Println(123123,j)
	(*j)[1]=123

	fmt.Println(j)

}