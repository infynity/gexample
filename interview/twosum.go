package main

import "fmt"

func main(){
	raw:=[4]int{2,3,7,9}

	target := 11

	fmt.Println(twosum(raw,target))
}

func twosum(raw [4]int,target int)[2]int{

	var arr [2]int

	hmap := make(map[int]int)

	for i:=0;i<len(raw);i++{
		complement:= target - raw[i]
		fmt.Println(raw[i],i)

		if val,ok:=hmap[complement];ok{

			arr[0]=i
			arr[1]=val
			break
		}
		hmap[raw[i]]=i

	}



	return arr

}
