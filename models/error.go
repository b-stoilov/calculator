package models

type Error struct {
	Expression string
	URL        string
	Frequency  int `default:"1"`
	Type       string
}
