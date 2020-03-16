package main

import "fmt"

func structTest() {
	//第一种方式
	stu := &Student{
		name: "Tom",
	}
	msg := stu.hello("jack")
	fmt.Printf(msg)
	//第二种方式
	stu2 := new(Student)
	fmt.Printf(stu2.hello("Sam"))

}

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {

	return fmt.Sprintf("hello %s i am %s", person, stu.name)
}
