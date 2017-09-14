package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}

}

func main() {
        //go keword calls the say function in separate task ...and this running say function will call as goroutine
	go say("Hello make Tech easier")
	//say("Hello make Tech easier")
	fmt.Println("sleep a little....")
        time.Sleep(10*time.Second)

}
