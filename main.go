package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}
	filepath := args[1]
	levels, err := readLevelStatsFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", levels)
}
