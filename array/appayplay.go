package main

import (
	"fmt"
)

func main() {
	score := [5]int{1: 1, 3: 5}
	fmt.Printf("%d\n", score[3])
	fmt.Printf("%d\n", score[2])

	score2 := [5]*int{1: new(int)}
	*score2[1] = 10
	fmt.Printf("%d\n", *score2[1])

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slice[2:3:5]
	slice2[0] = 9
	fmt.Println(slice[2])
	fmt.Println(slice2[0])
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	slice2 = append(slice2, 99)
	slice2 = append(slice2, 100)
	slice2 = append(slice2, 101)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slicea := []int{1, 2, 3, 4}
	sliceb := slicea[1:2:2]
	sliceb = append(sliceb, 10)
	sliceb = append(sliceb, 11)
	sliceb = append(sliceb, 12)
	sliceb = append(sliceb, 13)
	fmt.Println(slicea)
	fmt.Println(sliceb)

	fmt.Printf("%v\n", append(slicea, sliceb...))
	rangeSlice(&slicea)
	rangeSlice2(sliceb)
}

func rangeSlice(slice *[]int) {
	//函数间传递切片，并不复制底层数组，只复制切片本身，所以没必要这么搞
	for index, value := range *slice {
		fmt.Printf("index %d, value %d ,value addr %X, element addr %X\n", index, value, &value,
			&((*slice)[index]))
	}
}

func rangeSlice2(slice []int) {
	for _, value := range slice {
		fmt.Println(value)
	}
}
