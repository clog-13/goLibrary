package model

import "time"

type Feedback struct {
	FeedbackType string
	UserId       string
	ItemId       string
	Timestamp    time.Time
}
