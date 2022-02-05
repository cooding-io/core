package db

import (
	// postgres "github.com/lib/pq"
	"gorm.io/gorm"
)

// //Auth Auth from system
type Oauth2 struct {
	gorm.Model
	Name        string
	Description string
	ClientID    string
	SecretID    string
	Callback    string
	Enable      string
}

// type UsersApps struct {
// 	gorm.Model
// 	UserID int `sql:"type:int REFERENCES users(id)"`
// 	AppID  int `sql:"type:int REFERENCES apps(id)"`
// 	Enable bool
// 	Data   postgres.Jsonb
// }
