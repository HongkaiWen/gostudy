package main

import (
	"flag"
	"fmt"
	"mymath"
	"time"
)

func main() {
	var name_ = flag.String("name_", "xiaoming", "please input your name")
	flag.Parse()
	fmt.Println(*name_)

	fmt.Println("------------------")

	a, err := mymath.Add(1, 2)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err)
	}

	fmt.Println("------------------")

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
	fmt.Println("------------------")

	s := Student{Name: "xiaoming"}
	s.Name = "xiao"
	fmt.Println(s.Name)

	fmt.Println("------------------")

	fmt.Println("------------------")

	c := make(chan int, 10)
	defer close(c)
	go putChan(c)
	getChan(c)

}

func putChan(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}

}

func getChan(channel chan int) {

	for {
		timeout := make(chan bool, 1)
		defer close(timeout)
		go func() {
			time.Sleep(1e9) // 等待1秒钟
			timeout <- true
		}()

		select {
		case <-channel:
			result := <-channel
			fmt.Println(result)
		case <-timeout:
			fmt.Println("time out")
			break
		}

	}

}

type Student struct {
	Name string
}
