// auth: kunlun
// date: 2019-02-23
// description:  二分查找
package main

import "fmt"

func main() {

	l := make([]int, 100)
	for i := 1; i <= 100; i++ {
		l[i-1] = i
	}

	k := 77
	result, index, count := binary_search(l, k)
	fmt.Printf("search key: %v, result: %v, index: %v, count: %v\n", k, result, index, count)

}

// 二分法查找算法，需要保证输入是一个有序的元素列表
// 二分法时间复杂度为  O(log n)
// 算法步骤为：
//   1. 找到输入元素列表的起始索引
//   2. 找到列表的中间位置
//   3. 判断目标与中间位置数据的大小关系

func binary_search(l []int, k int) (result bool, mid int, count int) {
	start, end := 0, len(l)-1
	mid = (start + end) / 2
	for count = 1; count <= len(l); count++ {
		if k == l[mid] {
			result = true
			return
		} else if k > l[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
		mid = start + (end-start)/2
		fmt.Printf("start: %v, end: %v, middle: %v\n", start, end, mid)
	}
	return
}
