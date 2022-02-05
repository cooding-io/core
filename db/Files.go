package db

import "gorm.io/gorm"

//Files Files upload from system
type File struct {
	gorm.Model
	Name    string
	URL     string
	Enable  string
	OwnerID int
	Owner   User
}
