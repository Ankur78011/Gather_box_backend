package models

type Resturant struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Img_link    string `json:"img_link"`
	Description string `json:"desc"`
}

type NewResturant struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Img_link    string `json:"img_link" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateResturant struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Img_link    string `json:"img_link"`
	Description string `json:"description"`
}
