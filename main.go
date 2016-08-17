package main

import (
	"fmt"

	"github.com/mkappus/reserve/src/users"
)

func main() {
	us, err := users.NewStore("samp")
	if err != nil {
		panic(err)
	}
	fmt.Printf("US: %v", us)
}
