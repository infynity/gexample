package main

import "fmt"

func heapify(tree []int,n int,i int){

	if i>=n{
		return
	}
	c1 := i*2+1
	c2 := i*2+2
	max:=i
	if c1<n && tree[c1]>tree[max]{
		max=c1
	}
	if c2<n && tree[c2]>tree[max]{
		max=c2
	}
	if max!=i{
		tree[max],tree[i]=tree[i],tree[max]
		heapify(tree,n,max)
	}



}






func heapSorting(tree []int,n int){
	//构造大顶堆
	last:=n-1
	parent := (last-1)/2
	for i:=parent;i>=0;i--{
		heapify(tree,n,i)
	}
//------------
//	heapify(tree,n,0)

	fmt.Println(tree,"====",5/2)
	//return
//-----------


	for i:=n-1;i>=0;i--{
		//弹出堆顶最大值  最末节点放在堆顶

		tree[0],tree[i]=tree[i],tree[0]

		heapify(tree,i,0)//继续构造大顶堆

	}
}



func main(){

	sd :=[]int{4,10,3,5,1,2}
	sd =[]int{10,5,3,6,7,8}
	heapSorting(sd,6)

	fmt.Println(sd)
}