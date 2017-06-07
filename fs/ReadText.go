package main
import (
	"io/ioutil"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("e:/temp/tmp");
	if err != nil {
		fmt.Print(err.Error())
	}
	result := string(content)
	fmt.Print(result)
}
