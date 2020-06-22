package main

import (
	"fmt"
	"math/cmplx"
)

var (
	DeleteStatus bool       = false
	MaxInt       uint64     = 1<<64 - 1
	z            complex128 = cmplx.Sqrt(-5 + 12i)
)

// %v表示获取变量值
func main() {
	const boolStr = bool(false)
	const str = "hello world"
	var zhanghuan int
	var zhangxiang string
	var lisi bool
	fmt.Printf("type:\t %T  value:\t %v \n", false, false)
	fmt.Printf("zhanghuan value %v \t zhangxiang value %v lisi value %v \n", zhanghuan, zhangxiang, lisi)
	fmt.Printf("MaxInt %v", MaxInt)
	//Go 类型转换测试
	var i int = 32
	var f float64 = float64(i)
	var j uint64 = uint64(f)
	//简写
	q := 44
	w := float64(q)
	e := uint64(w)
	fmt.Printf("i value: %v\t f value: %v\t j value: %v\n", i, f, j)
	fmt.Printf("i value: %v\t f value: %v\t j value: %v\n", q, w, e)
}
