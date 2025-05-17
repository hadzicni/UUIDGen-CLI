package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	n := flag.Int("n", 1, "Number of UUIDs to generate")
	flag.Parse()

	if *n <= 0 {
		fmt.Println("n must be >= 1")
		os.Exit(1)
	}

	for i := 0; i < *n; i++ {
		fmt.Println(uuid.New())
	}
}
