package main

import "fmt"

func mapValue() {
	m1 := make(map[string]int) //仅声明
	m2 := map[string]string{
		"name": "zhang san",
		"age":  "23",
	}
	//赋值
	m1["Tom"] = 1
	//输出
	fmt.Println(m1, m2)
	m2["age"] = "33"
	fmt.Printf("修改后的m2 age值为 %s", m2)
}

/*
func main() {
	mapValue()
}
*/
