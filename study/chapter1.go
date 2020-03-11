package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main() {
	//声明方式
	/*var a int = 1
	var b int
	var c =1

	//简写
	b := 1
	var strrr = "hello word"

	//简单类型
	var e int8 =10
	var f byte ='a'
	var g float32 = 12.1
	var msg =  "hello word"
	ok := false*/
	ok := true
	fmt.Println(ok)
	//字符串
	var str1 = "Golang"
	str2 := "字符串2"
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2):", len(str2)) //一个字符串占3个字节  英文占1个字节  所以str2的长度为10
	/**
	golang 中字符串底层是通过byte数组实现的,所以直接求len 实际是在按字节长度计算,所以一个汉字占3个字节 算3个长度
	*/
	fmt.Println("处理字符串的正确方式为转换为rune数组：\n") //转换为数组类型 []rune 后，字符串中的每个字符 无论占多少字节都用 int32来表示
	str3 := "Go语音"
	runArry := []rune(str3)
	fmt.Println(len(str3)) //2+3*2 =8 //不能直接求 需要转换 两种方式
	fmt.Println(reflect.TypeOf(runArry[2]).Kind())
	fmt.Println(runArry[2], string(runArry[2]))
	fmt.Println("len(runeArr)", len(runArry)) //第一种获取字符串长度的方式

	runArryNew := utf8.RuneCountInString(str3) //第二种获取字符串长度的方式
	fmt.Println(runArryNew)

}
