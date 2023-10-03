package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
)

const ProfileRoot = "/Library/Application Support/Orion"
var uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

func searchBookmarks() error {
	showUpdateStatus()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var profiles []string
	entries, err := os.ReadDir(filepath.Join(home, ProfileRoot))
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == "Defaults" || uuidRegex.Match([]byte(entry.Name())) {
			profiles = append(profiles, entry.Name())
		}
	}

	var urls []string

	for _, profile := range profiles {
		plistPath := filepath.Join(home, ProfileRoot, profile, "favourites.plist")
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

			if slices.Contains(urls, item.Url) {
				continue
			}
			urls = append(urls, item.Url)

			wf.NewItem(item.Title).
				Valid(true).
				UID(id).
				Subtitle(item.Url).
				Match(match).
				Arg(item.Url)
		}

	}

	wf.WarnEmpty("No matching bookmarks found", "Try another search?")
	wf.SendFeedback()
	return nil
}

