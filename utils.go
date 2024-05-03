package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
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

func getSongUrl(levelId int) (string, error) {
	url := fmt.Sprintf("https://gdbrowser.com/%d", levelId)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", errors.New("request not ok")
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return "", err
	}

	a := querySelector(doc, `a[class="songLink"][href*=audio]`)
	if a == nil {
		return "", errors.New("no se encontró el elemento <a> de la canción")
	}

	href := ""
	for _, v := range a.Attr {
		if v.Key == "href" {
			href = v.Val
		}
	}

	if href == "" {
		return "", errors.New("no se encontró la url")
	}

	return href, nil
}

func querySelector(n *html.Node, query string) *html.Node {
	sel, err := cascadia.Parse(query)
	if err != nil {
		return &html.Node{}
	}
	return cascadia.Query(n, sel)
}
