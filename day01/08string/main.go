package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1, s2)
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	s3 := `第一行 \r \n \t \' \" \\
第二行
第三行
`
	fmt.Println(s3)

	fmt.Println(len(s1))
	fmt.Println(len(s2))

	fmt.Println(s1 + s2)

	s4 := fmt.Sprintf("%s is %s years old.\n", s1, s2)
	fmt.Println(s4)

	s5 := strings.Split(s4, "hello")
	fmt.Println(s5)

	isCon := strings.Contains(s4, "hello")
	fmt.Println(isCon)

	// strings.HasPrefix,strings.HasSuffix
	isHasPrefix := strings.HasPrefix(s4, "hello")
	fmt.Println(isHasPrefix)

	isHasSuffix := strings.HasSuffix(s1, "o")
	fmt.Println(isHasSuffix)

	// strings.Index(),strings.LastIndex()
	isIndex := strings.Index(s1, "l")
	fmt.Println(isIndex)

	lastIndex := strings.LastIndex(s1, "l")
	fmt.Println(lastIndex)

	// strings.Join(a[]string, sep string)

	// array of strings.
	str := []string{"Geeks", "For", "Geeks"}
	str2 := []string{"su", "jian", "hui"}
	// joining the string by separator
	fmt.Println(strings.Join(str, "-"))
	fmt.Println(strings.Join(str2, "-----"))
}
