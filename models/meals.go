package models

type MealInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Img_link       string `json:"img_link"`
	Resturant_name string `json:"resturant_name"`
	Prep_time      string `json:"prep_time"`
	Description    string `json:"desc"`
	Price          int    `json:"price"`
}

type NewMeal struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"desc" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Img_link    string `json:"img_link" binding:"required"`
	Prep_time   string `json:"prep_time" binding:"required"`
	Res_Id      int    `json:"res_id"`
}
