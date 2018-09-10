package main

import (
	"os"
	"fmt"
	"strconv"
)

func sum(vals ...interface{}) float64 {
	var res float64 = 0
	for _, val := range vals {
		switch val.(type) {
		case int:
			res += float64(val.(int))
		case int64:
			res += float64(val.(int64))
		case uint8:
			res += float64(val.(uint8))
		case string:
			a, err := strconv.ParseFloat(val.(string), 64)
			if err != nil {
				panic(err)
			}
			res += a
		default:
			fmt.Printf("Unsupported type %T. Ignoring.\n", val)
		}
	}
	return res
}

func main() {
	var a uint8 = 2
	var b int = 38
	var c string = "9.44"
	fmt.Printf("The result is %f\n", sum(a, b, c))
	fmt.Println(os.Getenv("GOROOT"))
}
