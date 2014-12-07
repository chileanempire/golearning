package main

import (
	"fmt"
	"time"

	"github.com/chileanempire/golearning/helloutils"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Current time is", time.Now())
	if err := helloutils.XD(11); err != nil {
		panic(err)
	}
}
