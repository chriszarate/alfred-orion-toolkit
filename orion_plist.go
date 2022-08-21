package main

import (
	"os"

  "howett.net/plist"
)

type BookmarkList map[string]Bookmark

type Bookmark struct {
	Keyword string `plist:"keyword"`
	Title string `plist:"title"`
	Type string `plist:"type"`
	Url string `plist:"url"`
}

type ReadingList struct {
	Id string `plist:"id"`
	Description string `plist:"description"`
	Title string `plist:"title"`
	Type string `plist:"type"`
	Url ReadingListUrl `plist:"url"`
}

type ReadingListUrl struct {
	Url string `plist:"relative"`
}

func ReadPlist[P BookmarkList | []ReadingList](filename string) (*P, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	d := plist.NewDecoder(file)
	var output P
	err = d.Decode(&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}
