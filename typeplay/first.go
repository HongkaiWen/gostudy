package main

import "fmt"

func main() {
	var s1 student
	fmt.Println(s1.name)

	s2 := student{
		name:    "jy",
		age:     3,
		address: "sz", //好奇葩，这块必须有逗号
	}

	fmt.Println(s2.name)
}

type student struct {
	name    string
	age     int
	address string
}
