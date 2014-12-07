package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/chileanempire/golearning/helloutils"
)

func main() {
	count := 5
	if len(os.Args) == 2 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			count = n
		}
	}
	fmt.Println("Hello, world!")
	fmt.Println("Current time is", time.Now())
	if err := helloutils.XD(count); err != nil {
		panic(err)
	}
}
