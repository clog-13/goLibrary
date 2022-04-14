package model

import "time"

type Item struct {
	ItemId     string
	IsHidden   bool
	Categories []string
	Timestamp  time.Time
	Labels     []string
	Comment    string
}

type PreItem struct {
	ItemId string
	BestId string
	Rate   string
}
