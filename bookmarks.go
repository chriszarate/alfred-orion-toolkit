package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const BookmarkPListPath = "/Library/Application Support/Orion/Defaults/favourites.plist"

func searchBookmarks() error {
	showUpdateStatus()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	plistPath := filepath.Join(home, BookmarkPListPath)

	b, err := ReadPlist[BookmarkList](plistPath)
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

	wf.WarnEmpty("No matching bookmarks found", "Try another search?")
	wf.SendFeedback()
	return nil
}
