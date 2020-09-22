package main

import (
	"context"
	"fmt"
)

func main(){


	arr :=[4]int8{2,4,6,8}

	fmt.Println(&arr[1],&arr[2])
	//0xc0000180d5 0xc0000180d6
	arr2 :=[4]int32{2,4,6,8}

	fmt.Println(&arr2[1],&arr2[2])
	//0xc000078034 0xc000078038
	fmt.Println()

	return
	slc := arr[1:3]
	slc =append(slc,12)
	//slc =append(slc,12)
	fmt.Println(slc)
	fmt.Println(&arr[1],&slc[0],cap(slc),len(slc))
	slc[1]=111
	fmt.Println(slc,arr)


	//ar := [3]int8{1,2,3}
	ar := [3]string{"abcde","山东省健康垃圾袋克拉斯担惊受恐接口的啊SDK阿拉善京东阿水","abcde"}

	fmt.Printf("%p\n",&ar)
	fmt.Println(&ar,&ar[0],&ar[1],&ar[2])


	var (
		cancelFunc context.CancelFunc
		cancelCtx context.Context

	)

	//cancelCtx, cancelFunc = context.WithCancel(context.TODO())


	fmt.Println(cancelCtx,cancelFunc,cancelFunc==nil)

}
