package models
import "time"
type Book struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Author string `json :"author"`
	CreatedAt time.Time `json :"created_at"`
	UpdatedAt time.Time `json :"updated_at"`
}




