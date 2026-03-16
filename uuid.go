package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1, _ := uuid.NewRandom()
	fmt.Printf("UUIDv4: %s\n", u1.String())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, _ := uuid.NewRandom()
	fmt.Printf("UUIDv4: %s\n", u2)
	u4 := uuid.New()
	fmt.Printf("UUIDv4: %s\n", u4)
	u5 := uuid.New()
	fmt.Printf("UUIDv4: %s\n", u5)
	// uuid  v1
	u3, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("UUID v1 %s\n", u3)
}
