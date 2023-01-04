package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1235123
	//整型转换字符串
	str := strconv.Itoa(num)
	// fmt.Printf("type:%T value is %s", str, str)

	strrune := []byte(str)
	for i, v := range strrune {
		count := 0
		for _, v2 := range strrune {
			if v == v2 {
				count++
			}
		}
		if count == 1 {
			fmt.Printf("出现一次的数字 %s", string(strrune[i]))
		}
	}
}
