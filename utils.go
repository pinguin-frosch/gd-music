package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func readLevelStatsFile(filepath string) ([]LevelInfo, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	levels := make([]LevelInfo, 0)
	err = json.Unmarshal(bytes, &levels)
	if err != nil {
		return nil, err
	}
	return levels, nil
}

func usage() {
	prog := os.Args[0]
	fmt.Printf("Usage: %v FILE\n\n", prog)
	fmt.Printf("FILE: Filepath of the file with the levels info in json format\n")
	os.Exit(1)
}

func getSongInfoOnline(levelId int) (int, error) {
	// TODO: usar gdbrowser.com para obtener esta información
	return levelId, errors.New("not implemented")
}
