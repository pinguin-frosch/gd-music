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

	songs := make(map[int]string)
	for _, l := range levels {
		if l.Type != "Online" {
			continue
		}

		songId := l.CustomSongId
		if songId == 0 {
			songId, err = getSongInfoOnline(l.Id)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		songs[songId] = fmt.Sprintf("https://www.newgrounds.com/audio/listen/%d", songId)
	}

	for _, s := range songs {
		fmt.Println(s)
	}
}
