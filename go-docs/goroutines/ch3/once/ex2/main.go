package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count += 1
	}
	decrement := func() {
		count -= 1
	}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)

	var onceA, onceB sync.Once
	var initB func()
	initA := func() {
		onceB.Do(initB)
	}
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
}
