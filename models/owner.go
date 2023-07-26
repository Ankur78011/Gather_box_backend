package models

type PostOwner struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Password string `json:"password"`
}
