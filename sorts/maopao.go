package sorts

import "fmt"

func Maopao(){
//3  2  1
//  1  2
//2  1  3
//  1


	//6+5+4+3+2+1

	toSort := []int{10,9,8,7,6,5}
	jkl := 0

	//todo 冒泡 每次找出最大的那个
	//todo N个数字要排序完成，总共进行N-1趟排序，第i趟的排序次数为(N-i)次
	lens :=len(toSort)-1   //5
	for i:=0;i<lens;i++{

		//for j:=i;j<lens;j++{  wrong
		//for j:=0;j<lens;j++{
		//for j:=0;j<i;j++{  big wrong
		for j:=0;j<lens-i;j++{

			jkl++
			if toSort[j]>toSort[j+1]{
				toSort[j],toSort[j+1] = toSort[j+1],toSort[j]
			}

		}
	}

	fmt.Println(toSort,jkl)

}
