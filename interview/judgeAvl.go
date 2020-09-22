package main

import (
	"fmt"
	"math"
)

//判断平衡二叉树
type BTree struct {

	value int
	left *BTree
	right *BTree
}
func NewBTree (value int) *BTree {
	return &BTree{value:value}
}
func(this *BTree) ConnectLeft(treeOrValue interface{}) *BTree{
	if bt,ok:=treeOrValue.(*BTree);ok{
		this.left=bt
	}else if v,ok:=treeOrValue.(int);ok{
		this.left=NewBTree(v)
	}
	return this
}
func(this *BTree) ConnectRight(treeOrValue interface{}) *BTree{
	if bt,ok:=treeOrValue.(*BTree);ok{
		this.right=bt
	}else if v,ok:=treeOrValue.(int);ok{
		this.right=NewBTree(v)
	}
	return this

}

func(this *BTree) String()   {
	fmt.Printf("二叉树:值是%d\n",this.value)
	var leftValue interface{}
	if this.left!=nil{
		leftValue=this.left.value
	}
	var rightValue interface{}
	if this.right!=nil{
		rightValue=this.right.value
	}
	fmt.Printf("左节点:%v   右节点:%v \n",leftValue,rightValue)
}



func main(){

		root:=NewBTree(13)
		//         13
		//    11         12

	root.ConnectLeft(11).ConnectRight(12)


	root.right.ConnectLeft(222)
	//  9   10		222

		root.left.ConnectLeft(9).ConnectRight(10)
	//root.left.left.ConnectLeft(7)

	judgeBstTrue(root)
	return

	s := getheight(root)
	fmt.Println(s,checkBal(root))


}


func getheight(nd *BTree) int {

	if nd==nil{
		return 0
	}

	leftheight := getheight(nd.left)
	rightheight := getheight(nd.right)

	ht := int(math.Max(float64(leftheight),float64(rightheight)))
	//ht:=0
	//if leftheight>rightheight{
	//	ht=leftheight
	//}else{
	//	ht=rightheight
	//}


	return ht +1
}


func checkBal(nd *BTree) bool{
	if nd ==nil{
		return true
	}

	l:=getheight(nd.left)
	r:=getheight(nd.right)

	if math.Abs(float64(l-r))>1{
		return false
	}
	return 	checkBal(nd.left) && checkBal(nd.right)

}

//只能判断局部   todo 中序遍历保存到数组 判断严格升序排列
func judgeBst(node *BTree)bool{
	if node==nil{
		return true
	}

	if node.left !=nil && node.left.value<node.value  && node.right!=nil && node.right.value>node.value{
		return true
	}else{
		return false
	}


	return judgeBst(node.left)&&judgeBst(node.right)
}




var arr = make([]int,0)

func judgeBstTrue(node *BTree)bool{
	midTraverse(node)

	fmt.Println(arr)
	for i:=0;i<len(arr);i++{
		if arr[i+1]<arr[i]{
			return false
		}
	}
	return true
}

func midTraverse(node *BTree){

	if node==nil{
		return
	}
	midTraverse(node.left)
	arr = append(arr,node.value)
	midTraverse(node.right)

}