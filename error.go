package main

import "fmt"

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s", e.message)
}
