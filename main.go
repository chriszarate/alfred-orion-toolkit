package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"log"
)

var (
	repo  = "chriszarate/alfred-orion-toolkit"
	query string
	wf    *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func main() {
	wf.Run(run)
}

func run() {
	if err := checkForUpdate(); err != nil {
		log.Printf("Error starting update check: %s", err)
	}
	args := wf.Args()
	switch args[0] {
	case "bookmarks":
		if err := searchBookmarks(); err != nil {
			log.Fatalf("search failed with %s", err.Error())
		}
	case "history":
		if err := searchHistory(); err != nil {
			log.Fatalf("search failed with %s", err.Error())
		}
	}
}
