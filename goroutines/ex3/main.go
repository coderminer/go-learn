package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside Again.")
	runtime.Gosched() // 这个main方法结束之前，会执行携程
}
