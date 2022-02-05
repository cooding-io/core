package db

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Finish      *time.Time
	Start       *time.Time
	SectionID   int
	OwnerID     int
	Integrates  int
	Role        int
	Owner       User
	Section     Section
}

type TaskQuestion struct {
	gorm.Model
	Question    string
	Description string
	Point       int
	Type        int
	TaskID      int
	Task        Task
}

type TaskQuestionAnswer struct {
	gorm.Model
	Name           string
	Description    string
	TaskQuestionID int
	Point          int
	Review         string
	FileID         int
	File           File
	TaskQuestion   TaskQuestion
}

type Answer struct {
	gorm.Model
	OwnerID      int
	TaskID       int
	Calification int
	Task         Task
	Owner        User
}

type AnswerUserFriends struct {
	gorm.Model
	AnswerID     int
	Answer       Answer
	UserID       int
	User         User
	Calification int
}
