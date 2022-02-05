package main

import db ".."
import "time"
import "io/ioutil"
import "fmt"
import "encoding/json"

type JSONData struct {
	Data struct {
		AllUsers struct {
			Edges []struct {
				Node struct {
					ID        int       `json:"id"`
					AboutMe   string    `json:"aboutMe"`
					Address   string    `json:"address"`
					Birthday  time.Time `json:"birthday"`
					CanvasID  int       `json:"canvasId"`
					Cover     string    `json:"cover"`
					Domain    string    `json:"domain"`
					Email     string    `json:"email"`
					LastName  string    `json:"lastName"`
					MoodleUDP bool      `json:"moodleUdp"`
					Name      string    `json:"name"`
					NodeID    string    `json:"nodeId"`
					Password  string    `json:"password"`
					Phone     int       `json:"phone"`
					Picture   string    `json:"picture"`
					Public    bool      `json:"public"`
					Ready     bool      `json:"ready"`
					Role      string    `json:"role"`
					Rut       string    `json:"rut"`
					CreatedAt time.Time `json:"createdAt"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"allUsers"`
	} `json:"data"`
}

func main() {
	db.Connect()

	file, _ := ioutil.ReadFile("u.json")

	data := JSONData{}

	_ = json.Unmarshal([]byte(file), &data)

	for _, user := range data.Data.AllUsers.Edges {

		u := db.User{}
		u.AboutMe = user.Node.AboutMe
		u.Address = user.Node.Address
		//	u.Birthday = &user.Node.Birthday
		u.CanvasID = user.Node.CanvasID
		u.Cover = user.Node.Cover
		u.Domain = user.Node.Domain
		u.Email = user.Node.Email
		u.LastName = user.Node.LastName
		u.MoodleUDP = user.Node.MoodleUDP
		u.Name = user.Node.Name
		u.Password = user.Node.Password
		u.Phone = user.Node.Phone
		u.Picture = user.Node.Picture
		u.Public = user.Node.Public
		u.Ready = user.Node.Ready
		u.Role = user.Node.Role
		u.Rut = user.Node.Rut
		fmt.Println(u)
		db.DB.Create(&u)

	}

	// user := db.User{}
	// user.Name = "Manuel"
	// user.LastName = "Alba"
	// user.Email = "malba@mmae.cl"
	// user.Role = "Pro"

	// db.DB.Create(&user)
}
