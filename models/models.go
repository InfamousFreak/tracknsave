package models 

/*import (
	"gorm.io/gorm"
)*/

/*type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
}*/

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID 			int 
	Email 		string 
	Password 	string
}