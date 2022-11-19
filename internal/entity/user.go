package entity

import "time"

type User struct {
	ID            int       `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Workplace     string    `json:"workplace"`
	Organization  string    `json:"organization"`
	CoverImageURL string    `json:"coverImageUrl"`
	Players       int       `json:"players"`
	Plays         int       `json:"plays"`
	Kahoots       int       `json:"kahoots"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
}
