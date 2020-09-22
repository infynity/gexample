package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

func insert(head *Node,newnode *Node){

	temp:=head
	for {
		if temp.Next==nil{
			//尾部
			break
		}
		temp=temp.Next
	}
	temp.Next=newnode

}


func showNOde(head *Node){
	temp:=head
	for {
		fmt.Println(temp.Value)

		if temp.Next==nil{
			//尾部
			break
		}
		temp=temp.Next
	}

}
func main(){

	createCu :=createRecursion([]int{1,2,3,4,5,6})
	traverse(createCu)
	return



	l1:=&Node{Value: 1}

	l2:=&Node{Value: 2}
	l3:=&Node{Value: 3}
	insert(l1,l2)

	insert(l1,l3)



	//traverseDD(l1)
	traverse(l1)
	traverse(reverseList(l1))
	return
	//
	//showList(l1)
	//return
	showNOde(l1)
	//return
	l1.Next=l2

	l2.Next=l3

	//showList(l1)
}


func traverseDD(head *Node){
	for p := head; p != nil; p = p.Next{
		// 迭代访问 p.val
		fmt.Println(p.Value)
	}
}

func traverse(head *Node) {
	if head==nil{
		return
	}
	fmt.Println(head.Value)
	// 递归访问 head.val
	traverse(head.Next)

}

func reverseList(head *Node) *Node{
	cur :=head
	var pre,temp *Node

	for  {
		if cur!=nil{
			temp=cur.Next
			cur.Next=pre
			pre=cur
			cur =temp
		}else{
			break
		}
	}
	return pre


}

func createRecursion(arr []int)*Node{

	if len(arr)==0{
		return nil
	}
	firstNode := &Node{Value: arr[0]}
	headoflist := createRecursion(arr[1:])

	firstNode.Next=headoflist
	return firstNode
}


//假设 翻转好了。。。
func fanzhuanRecursion(head *Node)*Node{

	if head==nil{
		return nil
	}
	if head.Next==nil{
		return head
	}

	newhead := fanzhuanRecursion(head.Next)
	head.Next.Next=head
	head.Next=nil
	return  newhead
}


func fanzhuanK(head *Node,k int) *Node{
	node:=head
	for i:=0;i<k;i++{
		if node==nil{
			return head
		}
		node=node.Next
	}

	newhead := reverseListOffset(head,node) //k=2 node =3
	head.Next = fanzhuanK(node,k)


	return newhead

}


func reverseListOffset(head *Node,end *Node) *Node{
	cur :=head
	var pre,temp *Node

	for  cur!=end{
			temp=cur.Next
			cur.Next=pre
			pre=cur
			cur =temp

	}
	return pre


}
