package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
)

var (
	n = flag.Int("n", 1, "Number of UUIDs to be generated")
)

func main() {
	flag.Parse()
	for i := 0; i < *n; i++ {
		fmt.Println(uuid.New())
	}
}
