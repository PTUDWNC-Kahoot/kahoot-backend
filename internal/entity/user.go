package entity

import "gorm.io/gorm"

type User struct {
	ID            uint32  `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Name          string `json:"name"`
	Workplace     string `json:"workplace"`
	Organization  string `json:"organization"`
	CoverImageURL string `json:"coverImageUrl"`
	Players       int8   `json:"players"`
	Plays         int8   `json:"plays"`
	Kahoots       int8   `json:"kahoots"`
	gorm.Model
}
