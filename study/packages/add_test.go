package main

import (
	"fmt"
	"testing"
)

//Test+方法名大写首字母
func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 5 {
		t.Error("error test!!!!")
	}
	fmt.Println("测试结果为：", result)
}
