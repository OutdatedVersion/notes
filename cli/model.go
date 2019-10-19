package main

import "time"

// Channel defines the top level bucket for notes to be held in
type Channel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Notes []Note `json:"notes"`
}

// Note defines the primary item we store
type Note struct {
	Content   string     `json:"content"`
	Active    bool       `json:"active"`
	CreatedAt *time.Time `json:"created_at"`
}
