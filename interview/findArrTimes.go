package main

import (
	"fmt"
)

func main(){
//	fmt.Println(os.Args[0])
//return
	rawArr :=[7]int{1,2,1,1,3,1,1}

	res :=finding(rawArr,2,8,1)

	fmt.Println(res)
}


func finding(rawArr [7]int,start int,end int ,tofind int)  int{

	//rawArr :=[5]int{1,2,1,1,3}


	midMap := make(map[int][]int)


	for i:=0;i<=len(rawArr)-1;i++{
		if v,ok:=midMap[rawArr[i]];!ok{
			fmt.Println(v)
			midMap[rawArr[i]]=make([]int,0)
		}
		//slc:=midMap[rawArr[i]]
		//slc = append(slc,i)
		//midMap[rawArr[i]]=slc
		midMap[rawArr[i]]=append(midMap[rawArr[i]],i)

	}

	fmt.Println(midMap)

	if val,ok:=midMap[tofind];!ok{
		return 0
	}else{


		fmt.Println(val)

		fmt.Println(countLessthan(val,start))

		fmt.Println(countLessthan(val,end+1))
		//
		//res :=0
		//for i,v:=range val{
		//	if v<=end && v>=start{
		//		res++
		//	}else if v>end {
		//		break
		//	}
		//	fmt.Println(i,v)
		//}
		return 100

	}



}
//小于less的有几个数
func countLessthan(rawArr []int,less int) int {

	l:=0
	r:=len(rawArr)-1
	mostRight :=-1

	for l<=r{
		mid := l+(r-l)/2
		if rawArr[mid]<less{
			mostRight=mid
			l=mid+1
		}else{
			r=mid-1
		}
	}

	return mostRight+1
}
