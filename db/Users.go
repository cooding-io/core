package db

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

//User user from system
type User struct {
	gorm.Model
	Name           string
	LastName       string
	Birthday       *time.Time
	Email          string
	Role           string
	Phone          int
	Address        string
	Domain         string
	Password       string
	AboutMe        string
	Ready          bool
	Public         bool
	Picture        string
	Cover          string
	Rut            string
	MoodleUDP      bool
	CanvasID       int
	Logins         datatypes.JSON
	AdditionalInfo datatypes.JSON
	Passwords      datatypes.JSON
}
