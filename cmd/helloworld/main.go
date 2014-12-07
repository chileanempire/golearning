package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Current time is", time.Now())
	XD(10)
}

// XD prints "XD" string n times.
func XD(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("XD")
	}
}
