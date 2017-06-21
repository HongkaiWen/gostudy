package main

import "fmt"

func main() {
	//init map
	mapPlay := map[string]string{"abc": "123", "def": "456"}

	//range map
	rangeMap(mapPlay)

	//get value
	fmt.Println(getValue("abc", mapPlay))
	fmt.Println(getValue("abc2", mapPlay))

	//delete value
	removeValue("abc", mapPlay)
	fmt.Println("after delete abc key ----")
	rangeMap(mapPlay)
}

func rangeMap(mapPlay map[string]string) {
	for key, value := range mapPlay {
		fmt.Println("key: " + key + ", vlaue: " + value)
	}
}

func getValue(key string, mapPlay map[string]string) string {
	value, _ := mapPlay[key]
	return value
}

func removeValue(key string, mapPlay map[string]string) {
	delete(mapPlay, key)
}
