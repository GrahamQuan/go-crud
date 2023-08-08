package models

import "gorm.io/gorm"

// must use caps letters to export Post
// same as "keys" of struct
type Post struct {
	gorm.Model
	Title string
	Body  string
}
