package db

import (
	"time"

	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Code         string
	InternalCode string
	Semester     int
	Year         int
	Section      int
	Start        *time.Time
	Finish       *time.Time
	Enable       bool
	Public       bool
	Scrapped     bool
	// Custom       string
	OwnerID  int
	Owner    User
	CourseID int
	Course   Course
}

type UsersSections struct {
	gorm.Model
	UserID    int
	User      User
	SectionID int
	Section   Section
	Role      int
}

type SectionFiles struct {
	gorm.Model
	FileID    int
	SectionID int
	Role      int
	File      File
	Section   Section
}
