package models

type CartItem struct {
	Date    string `json:"date"`
	Image   string `json:"image"`
	Name    string `json:"name"`
	Number  int    `json:"number"`
	Price   int    `json:"price"`
	Resname string `json:"resname"`
}
