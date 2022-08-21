package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const ReadingListPListPath = "/Library/Application Support/Orion/Defaults/reading_list.plist"

func searchReadingList() error {
	showUpdateStatus()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	plistPath := filepath.Join(home, ReadingListPListPath)

	r, err := ReadPlist[[]ReadingList](plistPath)
	if err != nil {
		return err
	}

	for _, item := range *r {
		match := fmt.Sprintf("%s %s", item.Title, item.Description)

		wf.NewItem(item.Title).
			Valid(true).
			UID(item.Id).
			Subtitle(item.Url.Url).
			Match(match).
			Arg(item.Url.Url)
	}

	wf.WarnEmpty("No matching reading list items found", "Try another search?")
	wf.SendFeedback()
	return nil
}
