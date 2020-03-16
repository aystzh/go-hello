package main

import "fmt"

func loopTests() {
	ifElse()
	switchTest()
	loopTest()
}

func ifElse() {
	age := 19
	if age < 19 {
		fmt.Println("kid")
	} else {
		fmt.Println("Adult")
	}
	//简写
	if sum := 20; sum < 20 {
		fmt.Println("small")
	} else {
		fmt.Println("big")
	}
}

func switchTest() {
	type Gender int8 //type定义了一种新的类型 Gender
	//类似java中的枚举类型
	const (
		MALE   Gender = 1
		FEMALE Gender = 2
	)
	gender := FEMALE
	//默认不需要break 标签 默认会跳出switch  如果需要继续执行 用 fallthrough
	switch gender {
	case MALE:
		fmt.Println("male")
	case FEMALE:
		fmt.Println("female")
		fallthrough //继续往下执行 default
	default:
		fmt.Println("unknown")
	}
}

//for循环的用法跟其他语言类似
func loopTest() {
	//简单类型 遍历
	sum := 0
	for i := 0; i < 10; i++ {
		if sum == 50 {
			break
		}
		sum += i
		fmt.Println("sum :", sum)
	}

	//复杂类型遍历  数组 切片 字典 map
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	m2 := map[string]string{
		"Sam":  "Male",
		"Jack": "Female",
	}
	//map 遍历
	for key, value := range m2 {
		fmt.Println(key, value)
	}

}
