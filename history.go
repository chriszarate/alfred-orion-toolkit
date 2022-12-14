package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	DBPath = "/Library/Application Support/Orion/Defaults/history"
	QUERY  = `
SELECT history_items.id, title, url
FROM history_items
INNER JOIN visits
ON visits.history_item_id = history_items.id
WHERE url LIKE ? OR title LIKE ?
GROUP BY url
ORDER BY count(visits.visit_time) DESC
`
)

func searchHistory() error {
	showUpdateStatus()

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dbPath := filepath.Join(home, DBPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("%%%s%%", strings.Join(strings.Split(os.Args[3], " "), "%"))
	rows, err := db.Query(QUERY, query, query)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id int
		var title, url sql.NullString
		err := rows.Scan(&id, &title, &url)
		if err != nil {
			return err
		}
		if !title.Valid || len(title.String) == 0 {
			title.String = url.String
		}
		wf.NewItem(title.String).
			Valid(true).
			UID(strconv.Itoa(id)).
			Subtitle(url.String).
			Arg(url.String)
	}
	wf.WarnEmpty("No matching history found", "Try another search?")
	wf.SendFeedback()
	return nil
}
