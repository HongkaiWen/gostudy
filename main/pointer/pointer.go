package main

import "fmt"

func main() {
	//	fmt.Println("hi go")
	total := 0
	test1(&total)
	fmt.Println(total)
}

func test1(total *int) {
	*total += 1
}
