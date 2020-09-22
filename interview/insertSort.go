package main

import "fmt"

func insertion(arr *[7]int) {

	for i := 1; i < len(arr); i++ {

		fmt.Println("本次数组", arr)
		fmt.Println()

		idx := i - 1
		val := arr[i]

		for idx >= 0 && arr[idx] < val {

			arr[idx+1] = arr[idx]
			fmt.Println(arr, "---", i, "---", idx)
			idx--
		}

		//if idx + 1 !=i{  TODO 这个判断应该可以去掉
		arr[idx+1] = val

		fmt.Printf("第%d位数组下表改变为%d\n", idx+1, val)
		fmt.Println()

		//}

	}
}

