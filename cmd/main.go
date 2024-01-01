package main

import (
	"fmt"

	"github.com/pborman/uuid"
)

func main() {
	fmt.Println("Hello!")
	testuuid := uuid.NewRandow()
	fmt.Println(testuuid)
}
