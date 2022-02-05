package db

import (
	"gorm.io/gorm"
	"time"
)

//Institution Institution from system
type Institution struct {
	gorm.Model
	Name    string
	Enable  bool
	Public  bool
	Expired *time.Time
	Domain  string
	OwnerID int
	Owner   User
	AboutMe string
	Picture string
	Cover   string
}

type UsersInstitutions struct {
	gorm.Model
	UserID        int
	User          User
	InstitutionID int
	Institution   Institution
	Role          int
}
