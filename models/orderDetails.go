package models

type OrderDetails struct {
	Name         string `json:"name" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required"`
	Address      string `json:"address" binding:"required"`
	ZipCode      string `json:"zipcode" binding:"required"`
	Email        string `json:"email" binding:"required"`
}

type OrderInfo struct {
	OrderId       int    `json:"order_id"`
	CustomerName  string `json:"customer_name"`
	MobileNumber  string `json:"mobile_number"`
	Address       string `json:"address"`
	ZipCode       string `json:"zipcode"`
	Email         string `json:"email"`
	MealName      string `json:"MealName"`
	DeliveryDate  string `json:"delivery_date"`
	Quantity      int    `json:"quantity"`
	Amount        int    `json:"amount"`
	ResturantName string `json:"resturant_name"`
}
type OrderCart struct {
	MealName    string `json:"meal_name"`
	DeliverDate string `json:"delivery_date"`
	Quantity    int    `json:"quantity"`
	Amount      int    `json:"amount"`
	Image_Link  string `json:"img_link"`
}
type OrderList struct {
	OrderId      int         `json:"order_id"`
	CustomerName string      `json:"customer_name"`
	MobileNumber string      `json:"mobile_number"`
	Address      string      `json:"address"`
	ZipCode      string      `json:"zipcode"`
	Email        string      `json:"email"`
	Order        []OrderCart `json:"orders"`
	OrderStatus  string      `json:"order_status"`
}
