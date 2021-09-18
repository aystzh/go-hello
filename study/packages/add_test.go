package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//Test+方法名大写首字母
func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 5 {
		t.Error("error test!!!!")
	}
	fmt.Println("测试结果为：", result)
}

func TestSwitch1(t *testing.T) {
	fmt.Print("Go run On...")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s . \n", os)
	}
}

func TestSwitch2(t *testing.T) {
	fmt.Println("test....")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("too far")
	}
}
