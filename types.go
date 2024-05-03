package main

import (
	"encoding/json"
	"os"
)

type LevelInfo struct {
	Id           int    `json:"id"`
	Type         string `json:"type"`
	CustomSongId int    `json:"customSongID"`
}

type State struct {
	CheckedIds map[int]bool   `json:"checked_ids"`
	SongUrls   map[string]int `json:"song_urls"`
}

type SongInfo struct {
	Url   string
	Count int
}

func getState() *State {
	bytes, err := os.ReadFile("state.json")
	if err != nil {
		return newState()
	}
	var state State
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		return newState()
	}
	return &state
}

func newState() *State {
	state := State{}
	state.SongUrls = make(map[string]int)
	state.CheckedIds = make(map[int]bool)
	return &state
}

func (s *State) Save() error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = os.WriteFile("state.json", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
