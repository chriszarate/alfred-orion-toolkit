package main

import (
	"fmt"
	"os"
	"path/filepath"

  "howett.net/plist"
)

type BookmarkList map[string]Bookmark

type Bookmark struct {
	Keyword string `plist:"keyword"`
	Title string `plist:"title"`
	Type string `plist:"type"`
	Url string `plist:"url"`
}

const PListPath = "/Library/Application Support/Orion/Defaults/favourites.plist"

func Readfile(filename string) (*BookmarkList, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	d := plist.NewDecoder(file)
	var bookmarks BookmarkList
	err = d.Decode(&bookmarks)
	if err != nil {
		return nil, err
	}

	return &bookmarks, nil
}

func searchBookmarks() error {
	showUpdateStatus()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	plistPath := filepath.Join(home, PListPath)

	b, err := Readfile(plistPath)
	if err != nil {
		return err
	}

	for id, item := range *b {
		if item.Type != "bookmark" {
			continue
		}

		match := fmt.Sprintf("%s %s", item.Title, item.Url)
		if item.Keyword != "" {
			match = fmt.Sprintf("%s %s", item.Keyword, match)
		}

		wf.NewItem(item.Title).
			Valid(true).
			UID(id).
			Subtitle(item.Url).
			Match(match).
			Arg(item.Url)
	}

	wf.WarnEmpty("No matching history found", "Try another search?")
	wf.SendFeedback()
	return nil
}
