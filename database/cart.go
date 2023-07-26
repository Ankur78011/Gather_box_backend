package database

import (
	"fmt"
	"time"
)

func (s *Storage) CartPost(name string, date string, number int, price int, orderId int) {
	sql := `select id from meals where name=$1`

	row := s.Db.QueryRow(sql, name)
	var meal_id int

	row.Scan(&meal_id)
	fmt.Println(meal_id, "uu")
	amount := number * price
	layout := "2006 Monday, January 2"
	currentYear := time.Now().Year()
	inputString := fmt.Sprintf("%d %s", currentYear, date)
	parsedDate, err := time.Parse(layout, inputString)
	if err != nil {
		panic(err)
	}

	sql2 := `insert into  ordered_items (meal_id,quantity,delivery_date,amount,price,order_id) values ($1,$2,$3,$4,$5,$6)`
	_, err = s.Db.Exec(sql2, meal_id, number, parsedDate, amount, price, orderId)
	if err != nil {
		panic(err)
	}

	fmt.Println(meal_id)

}
