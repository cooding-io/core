package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var databaseURL string

func init() {
	databaseURL = os.Getenv("database")
	databaseURL = "host=35.230.69.66 port=5432 user=postgres dbname=cooding password='lHCsCtckLkPkysmP' sslmode=require"
}

func Connect() {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	DB = db
}

func Setup() {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})

	// db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
		// panic("failed to connect database")
	}
	//defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Institution{}, &Course{}, &File{}, &Section{}, &Task{})
	db.AutoMigrate(&UsersSections{}, &SectionFiles{}, &UsersInstitutions{})
	db.AutoMigrate(&TaskQuestion{}, &TaskQuestionAnswer{})
	db.AutoMigrate(&Answer{}, &TaskQuestionAnswer{}, &AnswerUserFriends{})
	db.AutoMigrate(&Oauth2{})
	// db.AutoMigrate(&Message{}, &MessageFiles{})

}
