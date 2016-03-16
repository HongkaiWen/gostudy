package main

import (
	"fmt"
)

func main() {
	fmt.Println("abc")
	name := "suzhou"
	//按位去字符
	fmt.Println(name[0])
	//len函数
	fmt.Println(len(name))
	//遍历字符串
	for index, char := range name {
		fmt.Println(index, char)
	}
}
