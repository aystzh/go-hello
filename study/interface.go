package main

import "fmt"

//Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可。
type Person interface {
	getName() string
	//getAge() int
}

type StudentNew struct {
	name string
	age  int
}

func (stu *StudentNew) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.gender
}

func interfaceTest() {
	//将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。
	var _ Person = (*StudentNew)(nil)
	var _ Person = (*Worker)(nil)

	//实例转换为接口
	var p Person = &StudentNew{
		name: "王五",
		age:  0,
	}
	fmt.Printf(p.getName())

	var wo Person = &Worker{
		name:   "工人",
		gender: "男",
	}
	fmt.Println(wo.getName())

	wk := wo.(*Worker) //接口转换为实例
	fmt.Println(wk.name)

	//空接口
	blankInterface()
}

func blankInterface() {
	//如果定义了一个没有任何方法的空接口，那么这个接口可以表示任意类型。
	m := make(map[string]interface{})
	m["name"] = "tome"
	m["age"] = 12
	m["array"] = [3]int{1, 2, 3} //任意的value  相当于java 的 String,Object
	fmt.Println("空接口里面存放的值为：", m)
}
