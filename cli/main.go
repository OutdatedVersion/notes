package main

import (
	"flag"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	channelFlag = flag.String("c", "@me", "name of the channel to save into")
)

func main() {
	flag.Parse()

	log.Printf("content: %s\n", flag.Args())
	log.Printf("channel: %s\n", *channelFlag)

	// TODO(ben): configurable database location
	// database, _ := sql.Open("sqlite3", "./notes.db")
	// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS notes (id BLOB PRIMARY KEY, content TEXT, active TINYINT, created_at DATETIME);")
	// statement.Exec()
}
