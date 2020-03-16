package main

import (
	"fmt"
)

func functionTest() {
	res := addFunc(1, 4)
	fmt.Printf("1+4的结果为res: %d", res)

	//多个参数
	quo, rem, str := div(4, 2)
	fmt.Println("多个参数返回结果： ", quo, rem, str)

}

func addFunc(num1 int, num2 int) int {
	return num1 + num2
}

//多个参数测试
func div(num1 int, num2 int) (int, int, string) {
	return num1 / num2, num1 % num2, "haha"
}
