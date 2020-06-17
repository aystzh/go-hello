package main

import (
	"fmt"
	"time"
)

func main() {
	//简单类型 遍历
	begin := time.Now().UnixNano()
	var arr []int
	const NUMBER int = 10000000
	for i := 0; i < NUMBER; i++ {
		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		bubbleSort(arr)
	}
	fmt.Println("耗时：\n", (time.Now().UnixNano()-begin)/1000, "ms")
}

//排序
func bubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				temp := arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
}
