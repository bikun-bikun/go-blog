package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string
	Profile Profile
}

type Profile struct {
	gorm.Model
	UserId       int
	Repositories []Repository
}

type Repository struct {
	gorm.Model
	ProfileId int
	Name      string
	Language  string
}
