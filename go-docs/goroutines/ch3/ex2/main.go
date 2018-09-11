package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)

	var wg1 sync.WaitGroup
	for _, val := range []string{"hello", "greetings", "good day"} {
		wg1.Add(1)
		go func(v string) {
			defer wg1.Done()
			fmt.Println(v)
		}(val)
	}
	wg1.Wait()
}
