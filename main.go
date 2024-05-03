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

	urls := make([]string, 0)
	for _, l := range levels {
		if l.Type != "Online" {
			continue
		}
		songUrl, err := getSongUrl(l.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		urls = append(urls, songUrl)
	}

	fmt.Println(urls)
}
