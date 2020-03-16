package main

import (
	"errors"
	"fmt"
	"os"
)

func errorTest() {
	//错误异常处理
	errorHanding()
	//自定义异常处理
	if err := errorNew(""); err != nil {
		fmt.Println(err)
	}
	//panic
	res := errorPanic(2)
	fmt.Println(res)
	//自定义返回异常信息
	var autoErrInfo = errorTryCatch(3)
	fmt.Println("自定义返回的异常结果为：", autoErrInfo)
}
func errorHanding() {
	files, err := os.Open("loop.go")
	if err != nil {
		fmt.Println(err) // no such file or directory
	}
	fmt.Println(&files)
}

//error.New 返回自定义异常
func errorNew(str string) error {
	if len(str) == 0 {
		return errors.New("字符串不能为空!")
	}
	return nil
}

//error为可以预知的异常  不可预知的异常成为 panic  类似于java的 RuntimeException
func errorPanic(index int) int {
	array := [3]int{1, 2, 3}
	return array[index]
}

// defer 和 recover 捕获异常并返回处理结果 类似于java的 try catch
//使用 defer 定义了异常处理的函数，在协程退出前，会执行完 defer 挂载的任务。因此如果触发了 panic，控制权就交给了 defer。
//在 defer 的处理逻辑中，使用 recover，使程序恢复正常，并且将返回值设置为 -1，在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0。
func errorTryCatch(index int) (res int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("some error has happened：", r)
			res = -1 //把返回值设为-1并返回 如果不设置则默认返回0
		}
	}()
	array := [3]int{1, 2, 3}
	return array[index]
}
