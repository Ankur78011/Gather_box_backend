package database

import (
	"fmt"

	"gatherbox.com/models"
)

func (s *Storage) GetMeals() []models.MealInfo {
	rows, err := s.Db.Query("select m.id,m.name,r.name,m.price,m.description,m.img_link,m.prep_time from meals as m inner join resturants as r on r.id=m.res_id ")
	if err != nil {
		panic(err)
	}
	var ListOfMeals []models.MealInfo
	for rows.Next() {
		SingleMealInfo := models.MealInfo{}
		err := rows.Scan(&SingleMealInfo.Id, &SingleMealInfo.Name, &SingleMealInfo.Resturant_name, &SingleMealInfo.Price, &SingleMealInfo.Description, &SingleMealInfo.Img_link, &SingleMealInfo.Prep_time)
		if err != nil {
			panic(err)
		}
		ListOfMeals = append(ListOfMeals, SingleMealInfo)
	}
	return ListOfMeals
}
func (s *Storage) GetMealsAsPerPrice(price int) []models.MealInfo {
	sql := `select m.id,m.name,r.name,m.price,m.description,m.img_link,m.prep_time from meals as m inner join resturants as r on r.id=m.res_id  where m.price<$1`
	rows, err := s.Db.Query(sql, price)
	if err != nil {
		panic(err)
	}
	var ListOfMeals []models.MealInfo
	for rows.Next() {
		SingleMealInfo := models.MealInfo{}
		err := rows.Scan(&SingleMealInfo.Id, &SingleMealInfo.Name, &SingleMealInfo.Resturant_name, &SingleMealInfo.Price, &SingleMealInfo.Description, &SingleMealInfo.Img_link, &SingleMealInfo.Prep_time)
		if err != nil {
			panic(err)
		}
		ListOfMeals = append(ListOfMeals, SingleMealInfo)
	}
	return ListOfMeals
}

func (s *Storage) CreateMeal(name string, desc string, price int, img_link string, prep_time string, res_id int) {
	sqlStatment := `Insert into meals (name,description,price,img_link,prep_time,res_id) values ($1,$2,$3,$4,$5,$6)`
	_, err := s.Db.Exec(sqlStatment, name, desc, price, img_link, prep_time, res_id)
	if err != nil {
		panic(err)
	}
}

func (s *Storage) GetResWiseMeal(id int) []models.MealsFromRes {

	sqlStatament := `select r.name,m.name,m.price,m.img_link,m.id from meals as m inner join resturants as r on r.id=m.res_id where r.id=$1`
	rows, err := s.Db.Query(sqlStatament, id)
	if err != nil {
		panic(err)
	}
	var ListOfMealFromRes []models.MealsFromRes
	for rows.Next() {
		SingleMealFromRes := models.MealsFromRes{}
		err := rows.Scan(&SingleMealFromRes.Res_name, &SingleMealFromRes.Name, &SingleMealFromRes.Price, &SingleMealFromRes.Img_link, &SingleMealFromRes.Id)
		if err != nil {
			panic(err)
		}
		ListOfMealFromRes = append(ListOfMealFromRes, SingleMealFromRes)
	}
	return ListOfMealFromRes
}

func (s *Storage) GetMealsResId(res_id int) []models.MealInfo {
	sql := `Select m.id,m.name,m.img_link,r.name,m.prep_time,m.description,m.price  from meals as m inner join resturants as r on r.id=m.res_id where r.id=$1`
	rows, err := s.Db.Query(sql, res_id)
	if err != nil {
		panic(err)
	}
	var listOfMeals []models.MealInfo
	for rows.Next() {
		var singleMeal models.MealInfo
		err := rows.Scan(&singleMeal.Id, &singleMeal.Name, &singleMeal.Img_link, &singleMeal.Resturant_name, &singleMeal.Prep_time, &singleMeal.Description, &singleMeal.Price)
		if err != nil {
			panic(err)
		}
		listOfMeals = append(listOfMeals, singleMeal)

	}
	return listOfMeals
}

func (s *Storage) CreateTheMeal(name string, imgLink string, desc string, prep_time string, price int, res_id int) {
	sql := `Insert into meals (name,description,price,img_link,prep_time,res_id) values($1,$2,$3,$4,$5,$6)`
	_, err := s.Db.Exec(sql, name, desc, price, imgLink, prep_time, res_id)
	if err != nil {
		panic(err)
	}

}

func (s *Storage) CheckThatOwnerDeleteItsMeal(ownerId int, meal_id int) bool {
	sql := `select id from meals  where res_id in(select id from resturants where owner_id=$1)`
	rows, err := s.Db.Query(sql, ownerId)
	if err != nil {
		fmt.Println(err)
		return false
	}
	for rows.Next() {
		var id int
		_ = rows.Scan(&id)
		if id == meal_id {
			return true
		}

	}
	return false

}

func (s *Storage) DelteFromMealTable(meal_id int) {
	sql := `Delete from meals where id=$1`
	_, err := s.Db.Exec(sql, meal_id)
	if err != nil {
		panic(err)
	}

}

func (s *Storage) UpdateMeals(name string, desc string, price int, prepTime string, imgLink string, meal_id int) {
	sql := `Update meals SET name=$1,description=$2,price=$3,prep_time=$4,img_link=$5 where id=$6`
	_, err := s.Db.Exec(sql, name, desc, price, prepTime, imgLink, meal_id)
	if err != nil {
		panic(err)
	}
}
