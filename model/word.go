package model

import "time"

type Word struct {
	Id        int       `json:"id" xorm:"pk autoincr"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	MediaId   string    `json:"media_id"`
	CreatedAt time.Time `json:"created_at" xorm:"<-"`
}
