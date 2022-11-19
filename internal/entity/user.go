package entity

import "time"

type User struct {
	Id              int
	Email           string
	Password        string
	Name            string
	Workplace       string
	Organization    string
	Cover_image_url string
	Players         int
	Plays           int
	Kahoots         int
	Created_at      time.Time
	Updated_at      time.Time
	Deleted_at      time.Time
}
