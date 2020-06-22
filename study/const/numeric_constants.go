package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Printf("small:%v \n", Small)
	fmt.Printf("needInt %v \n", needInt(Small))
	fmt.Printf("needInt %v \n", needFloat(Small))
	fmt.Printf("needInt %v \n", needFloat(Big))
	loop(0)
	fmt.Println(sqrt(-4))
	fmt.Println("power:", power(2, 3, 3))
}

func needInt(x int) (res int) {
	res = x*10 + 1
	return res
}

func needFloat(x float64) (res float64) {
	res = x * 0.1
	return res
}

func loop(num int) {
	for i := 0; i < 10; i++ {
		num += i
		fmt.Println(num)
	}

	k := 0
	for ; k < 100; k++ {
		num += k
		fmt.Println(num)
	}

	for k < 100 {
		num += k
		fmt.Println(num)
	}
}

func sqrt(i float64) string {
	if i < 0 {
		return sqrt(-i) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}

//开i的j次方  math.Pow(i,j)
func power(i, j, k float64) float64 {
	if w := math.Pow(i, j); w < k {
		return w
	} else {
		fmt.Printf("%g >= %g\n", w, k)
	}
	return k

}
