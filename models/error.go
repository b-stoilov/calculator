package models

type Error struct {
	Expression string
	URL        string
	Frequency  int
	Type       string
}
