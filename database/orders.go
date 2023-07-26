package database

import (
	"gatherbox.com/models"
)

func (s *Storage) InsertOrderDetails(name string, mobile string, address string, zipcode string, email string, user_id int) int {
	sql := `INSERT INTO order_details (name,mobile,address,zipcode,email,user_id,order_status) values($1,$2,$3,$4,$5,$6,'Pending') RETURNING id`
	var order_id int
	row := s.Db.QueryRow(sql, name, mobile, address, zipcode, email, user_id)
	err := row.Scan(&order_id)
	if err != nil {
		panic(err)
	}
	return order_id
}

func (s *Storage) DeleteFromOrder(meal_id int) {
	sql := `Delete From ordered_items where meal_id=$1`
	_, err := s.Db.Exec(sql, meal_id)
	if err != nil {
		panic(err)
	}

}
func (s *Storage) GetOrder(res_id int) []models.OrderInfo {
	sql := `select i.order_id,d.name,d.mobile,d.address,d.zipcode,d.email,m.name,i.delivery_date,i.quantity,i.amount,r.name
from meals as m
inner join ordered_items as i
on m.id=i.meal_id
inner join order_details as d
on d.id=i.order_id
inner join resturants as r
on r.id=m.res_id
where m.res_id=$1`

	rows, err := s.Db.Query(sql, res_id)
	if err != nil {
		panic(err)
	}
	var OrderList []models.OrderInfo
	for rows.Next() {
		var Order models.OrderInfo
		err := rows.Scan(&Order.OrderId, &Order.CustomerName, &Order.MobileNumber, &Order.Address, &Order.ZipCode, &Order.Email, &Order.MealName, &Order.DeliveryDate, &Order.Quantity, &Order.Amount, &Order.ResturantName)
		if err != nil {
			panic(err)
		}
		OrderList = append(OrderList, Order)
	}
	return OrderList

}

func (s *Storage) GetCart(res_id int) []models.OrderList {
	sql := `select i.order_id,d.name,d.mobile,d.address,d.zipcode,d.email,d.order_status
 from meals as m
 inner join ordered_items as i
 on m.id=i.meal_id
 inner join order_details as d
 on d.id=i.order_id
 inner join resturants as r
 on r.id=m.res_id
 where m.res_id=$1
 Group by i.order_id,d.name,d.mobile,d.address,d.zipcode,d.email,d.order_status`
	var OrderDet []models.OrderList
	rows, err := s.Db.Query(sql, res_id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var Ord models.OrderList
		err := rows.Scan(&Ord.OrderId, &Ord.CustomerName, &Ord.MobileNumber, &Ord.Address, &Ord.ZipCode, &Ord.Email, &Ord.OrderStatus)
		if err != nil {
			panic(err)
		}
		sql2 := `select m.name,i.delivery_date,i.quantity,i.amount,m.img_link
		from meals as m
		inner join ordered_items as i
		on m.id=i.meal_id
		inner join order_details as d
		on d.id=i.order_id
		inner join resturants as r
		on r.id=m.res_id
		where m.res_id=$1 and order_id=$2`
		rows, err := s.Db.Query(sql2, res_id, Ord.OrderId)
		var OrderItems []models.OrderCart
		for rows.Next() {
			var Ord2 models.OrderCart
			err := rows.Scan(&Ord2.MealName, &Ord2.DeliverDate, &Ord2.Quantity, &Ord2.Amount, &Ord2.Image_Link)
			if err != nil {
				panic(err)
			}
			OrderItems = append(OrderItems, Ord2)

		}
		Ord.Order = OrderItems
		OrderDet = append(OrderDet, Ord)
	}
	return OrderDet

}

func (s *Storage) ChangeStatus(order_id int) {
	sql := `UPDATE order_details set order_status='Delivered' Where id=$1`
	_, err := s.Db.Exec(sql, order_id)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) GetStatus(order_id int) string {
	sql := `select order_status from order_details where id=$1`
	row := s.Db.QueryRow(sql, order_id)
	var status string
	err := row.Scan(&status)
	if err != nil {
		panic(err)
	}
	return status
}

func (s *Storage) UpdateStatusToDelivered(order_id int) {
	sql := `Update order_details SET order_status='Delivered' Where id=$1`
	_, err := s.Db.Exec(sql, order_id)
	if err != nil {
		panic(err)
	}

}
func (s *Storage) UpdateStatusToCancelled(order_id int) {
	sql := `Update order_details SET order_status='Cancelled' Where id=$1`
	_, err := s.Db.Exec(sql, order_id)
	if err != nil {
		panic(err)
	}
}
