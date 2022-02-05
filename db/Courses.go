package db

import (
	"gorm.io/gorm"
)

//Courses Courses from system
type Course struct {
	gorm.Model
	Name         string
	Code         string
	InternalCode string
	Enable       bool
	Public       bool
	Scrapped     bool
	// Custom             string
	Description        string
	Video              string
	Picture            string
	OwnerInstitutionID int
	OwnerInstitution   Institution
}
