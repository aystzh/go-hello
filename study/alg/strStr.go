package main

import "fmt"

//source  123456     dest  45e
func strStr(sourceStr string, destStr string) int {
	if len(destStr) == 0 {
		return -1
	}
	var i, j int

	for i = 0; i < len(sourceStr)-len(destStr)+1; i++ {
		for j = 0; j < len(destStr); j++ {
			if sourceStr[i+j] != destStr[j] {
				break
			}
		}

		if len(destStr) == i {
			return i
		}
	}
	return 1
}

func main() {
	i := strStr("123456", "45e")
	fmt.Print("第一次出现的位置是:", i)
}
