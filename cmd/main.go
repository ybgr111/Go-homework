package main

import (
	"fmt"

	"github.com/pborman/uuid"
)

func main() {

	fmt.Println("Hello!")

	testuuid := uuid.NewRandom()
	fmt.Println(testuuid)
	fmt.Println(testuuid)
}
