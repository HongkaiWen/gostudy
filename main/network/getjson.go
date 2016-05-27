package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
*http://blog.csdn.net/kenkao/article/details/47857757
 */
func main() {
	resp, err := http.Get("http://")
	if err != nil {
		fmt.Println("network error")
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io error")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}
