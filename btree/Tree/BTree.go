package Tree

import (
	"fmt"
)

type BTree struct {
	Value int
	Left *BTree
	Right *BTree
}
func(this *BTree) String()   {
	  fmt.Printf("二叉树:值是%d\n",this.Value)
	  var leftValue interface{}
	  if this.Left!=nil{
	  	leftValue=this.Left.Value
	  }
	var rightValue interface{}
	if this.Right!=nil{
		rightValue=this.Right.Value
	}
	  fmt.Printf("左节点:%v   右节点:%v \n",leftValue,rightValue)
}  //以字符串的形式打印
func max(a int ,b int) int   {
	if a>=b{
		return a
	}
	return b
}
//获取层高
func(this *BTree) Level() int  {
	 if this==nil{
		return 0
	}
	return max(this.Left.Level(),this.Right.Level())+1
}

//先序遍历
func(this *BTree)Preorder()  {
	if this==nil{
		return
	}
	fmt.Printf("%d->",this.Value)
	//先判断左节点
	this.Left.Preorder()
	this.Right.Preorder()

}

func(this *BTree)Midorder()  {
	if this==nil{
		return
	}
	this.Left.Midorder()

	fmt.Printf("%d->",this.Value)
	//先判断左节点
	this.Right.Midorder()

}
func(this *BTree) ConnectLeft(treeOrValue interface{}) *BTree{
	if bt,ok:=treeOrValue.(*BTree);ok{
		this.Left=bt
	}else if v,ok:=treeOrValue.(int);ok{
		this.Left=NewBTree(v)
	}
	 return this
}
func(this *BTree) ConnectRight(treeOrValue interface{}) *BTree{
	if bt,ok:=treeOrValue.(*BTree);ok{
		this.Right=bt
	}else if v,ok:=treeOrValue.(int);ok{
		this.Right=NewBTree(v)
	}
	return this

}

func NewBTree (value int) *BTree {
	return &BTree{Value:value}
}



type BTrees []*BTree //二叉树集合类型
func(this BTrees) String(){
	for _,bt:=range this{
		bt.String()
	}
	fmt.Printf("当前一共有%d个节点",len(this))
}
func NewBTrees (values... int) BTrees{
	 btrees:=make(BTrees,len(values))
	 for index,v:=range values{
		 btrees[index]=NewBTree(v)
	 }
	 return btrees
}
