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

	state := getState()
	fmt.Printf("%+v\n", state)

	for _, l := range levels {
		fmt.Printf("%v", l.Id)
		if state.CheckedIds[l.Id] {
			fmt.Println(" skipped")
			continue
		}
		fmt.Println()

		if l.Type != "Online" {
			continue
		}
		songUrl, err := getSongUrl(l.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		state.CheckedIds[l.Id] = true
		_, ok := state.SongUrls[songUrl]
		if !ok {
			state.SongUrls[songUrl] = 0
		}
		state.SongUrls[songUrl] += 1

		state.Save()
	}

	fmt.Printf("%+v\n", state.SongUrls)
}
