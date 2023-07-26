package models

type User struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	UserName string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
}



