package entity

import "time"

type User struct {
	ID            uint8     `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Name          string    `json:"name"`
	Workplace     string    `json:"workplace"`
	Organization  string    `json:"organization"`
	CoverImageURL string    `json:"coverImageUrl"`
	Players       int8      `json:"players"`
	Plays         int8      `json:"plays"`
	Kahoots       int8      `json:"kahoots"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	DeletedAt     time.Time `json:"deletedAt"`
}
