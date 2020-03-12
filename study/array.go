package main

import "fmt"

func array() [5]int {
	//  var arr2 [5][5]int 二维数组
	//使用[] 索引/修改数组
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	return arr
}

func slice() {
	//slice1 := make([]float32, 0)//长度为0的切片
	slice2 := make([]float32, 3, 5) //长度为3容量为5的切片
	fmt.Println(len(slice2), cap(slice2))

	//切片使用 添加切片 切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4, 5)
	fmt.Println(len(slice2), cap(slice2), "slice : ", slice2)

	sub1 := slice2[3:] //[1 2 3 4 5]
	fmt.Println("sub1: ", sub1)
	sub2 := slice2[:3]
	fmt.Println("sub2:", sub2)
	sub3 := slice2[1:5]
	fmt.Println("sub3 :", sub3)
	combined := append(sub1, sub2...) //合并切片
	fmt.Println("合并后的切片为:", combined)

}

func main() {
	slice()
}
