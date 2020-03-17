package main

import "fmt"

//Go 语言也有 Public 和 Private 的概念，粒度是包。如果类型/接口/方法/函数/字段的首字母大写，则是 Public 的，
//对其他 package 可见，如果首字母小写，则是 Private 的，对其他 package 不可见。
func main() {
	res := add(1, 2)
	fmt.Println(res)
}
