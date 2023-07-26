package models

type MealsFromRes struct {
	Name     string `json:"name"`
	Res_name string `json:"res_name"`
	Img_link string `json:"img_link"`
	Price    int    `json:"price"`
	Id       int    `json:"id"`
}
