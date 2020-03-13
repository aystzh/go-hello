package main

import "fmt"

func pointerValue() {
	str := "hello Golang"
	var p *string = &str
	*p = "word"
	fmt.Println("指针修改过后的值为: ", str)
}

func main() {
	pointerValue()
	num := 100
	add(num)
	fmt.Println("值传递修改过后的值为：", num) //值传递不会修改值本身。函数内部会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值

	reAdd(&num)
	fmt.Println("引用传递修改过后的值为：", num) //参数使用指针 对参数的传递会影响到外部变量
}

func add(num int) {
	num += 1
}

func reAdd(num *int) {
	*num += 1
}
