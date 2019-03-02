package main

import (
	"fmt"
	"time"
        "math"
)
type error interface {
    Error() string
}
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() (float64, error) {
	return 0, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Sqrt(-2))
}
