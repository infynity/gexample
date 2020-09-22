package main


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

		// l >= r 表明本次分解任务完成, break
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
// 快速  平衡  堆排序  翻转