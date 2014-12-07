package helloutils

import (
	"errors"
	"fmt"
)

// XD prints "XD" string n times.
func XD(n int) error {
	if n > 10 {
		return errors.New("too many XD's to print")
	}
	for i := 0; i < n; i++ {
		fmt.Println("XD")
	}
	return nil
}
