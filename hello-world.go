package main

import "fmt"

func hello() int {
	fmt.Println("its hello function calling")
	return 10
}
func main() {
	fmt.Println("hello world")
	test("calling test function")
}
