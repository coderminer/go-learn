package main

import (
	"fmt"
	"os"
	"time"
)

/*
	channel 可以实现从一个协程发消息到另一个协程
*/

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}

func main() {
	done := time.After(30 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			fmt.Println(string(buf))
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}
