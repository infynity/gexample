package sorts

import (
	"fmt"
	"math/rand"
	"time"
)

//冒泡排序
func BubbleSort(arr *[8]int) {

	fmt.Println("排序前arr=", (*arr))

	jkl := 0
	//冒泡排序..一步一步推导出来的
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			jkl++
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				arr[j], arr[j+1] = arr[j+1], arr[j]

			}
		}

	}

	fmt.Println("排序后arr=", (*arr), jkl, "次")

}

//从da到dxiao排序  todo 选择排序
func choose_sort(arr *[7]int) {

	for j := 0; j < len(arr)-1; j++ {
		fmt.Println("本次数组", arr)

		maxval := arr[j]
		maxid := j
		for i := j + 1; i < len(arr); i++ {
			if arr[i] > maxval {
				maxval = arr[i]
				maxid = i
			}

		}

		arr[j], arr[maxid] = arr[maxid], arr[j]
		fmt.Println("第", j, "次排序完毕", arr)
		fmt.Println()

	}

}

//TODO 插入排序

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

func speed() {
	var arr [80000]int

	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(90000)
	}

	start := time.Now().Unix()

	// sort here

	end := time.Now().Unix()

	fmt.Println(start - end)
}

func main() {

	brr := [7]int{4, 6, 511, 5131, 10086, 33, 1957}

	fmt.Println("原数组", brr)
	insertion(&brr)

	fmt.Println("插入排序完成的结果", brr)

	//this is choose sorting
	//and the above one is insertion sorting

	//todo

	//please look down upon the description and see the running result

	arr := [7]int{4, 6, 511, 521, 10086, 33, 1957}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	//
	//arr[1]=arr[0]
	//fmt.Println(arr)
	//return
	choose_sort(&arr)

	fmt.Println("排序完成的结果", arr)

	fmt.Println("快速排序如下")

	//快速排序   递归  这个比较难以理解
	crr := [9]int{-1, -3233, 144, 123, 434, 3231, 0, -10, 23}
	quick_sort(0, 8, &crr)
	fmt.Println(crr)

	drr := [8]int{2, 3, 1, 5, 11, 144, 32, 12}
	BubbleSort(&drr)
	fmt.Println(drr)
}

//快速排序    有点像二叉树？

//todo 2019.11.01   18:19  没有理解 yet   这个是最快的 排序 八百万数据 才 3秒就排好
// todo  但是 耗费的资源比较多  开的栈很多 （递归）
// 时间复杂度
func quick_sort(left int, right int, array *[9]int) {

	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	//for 循环的目标是将比 pivot 小的数放到 左边 // 比 pivot 大的数放到 右边
	for l < r {
		//从 pivot 的左边找到大于等于 pivot 的值
		for array[l] < pivot {
			l++
		}
		//从 pivot 的右边边找到小于等于 pivot 的值
		for array[r] > pivot {
			r--
		}

		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换

		array[l], array[r] = array[r], array[l]

		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}

	// 如果
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		quick_sort(left, r, array)
	}

	// 向右递归
	if right > l {
		quick_sort(l, right, array)
	}
}
