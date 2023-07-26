package database

import (
	"gatherbox.com/models"
)

func (s *Storage) GetResturants() []models.Resturant {
	rows, err := s.Db.Query("Select id,name,address,img_link,description from resturants Order by id")
	if err != nil {
		panic(err)
	}
	var ResturantList []models.Resturant
	for rows.Next() {
		Single_Rest := models.Resturant{}
		err := rows.Scan(&Single_Rest.Id, &Single_Rest.Name, &Single_Rest.Address, &Single_Rest.Img_link, &Single_Rest.Description)
		if err != nil {
			panic(err)
		}
		ResturantList = append(ResturantList, Single_Rest)
	}
	return ResturantList
}

func (s *Storage) CreateResturant(owner_id int, name string, address string, img_link string, desc string) {
	sqlStatment := `INSERT INTO resturants (name,address,owner_id,img_link,description) values($1,$2,$3,$4,$5)`
	_, err := s.Db.Exec(sqlStatment, name, address, owner_id, img_link, desc)
	if err != nil {
		panic(err)
	}

}
func (m *Storage) GetRecentRes() []models.ResturantName {
	sql := `select name,id,img_link from resturants limit 4`
	var RecentList []models.ResturantName
	rows, err := m.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var NewRes models.ResturantName
		err := rows.Scan(&NewRes.Name, &NewRes.Id, &NewRes.Img_link)
		if err != nil {
			panic(err)
		}
		RecentList = append(RecentList, NewRes)
	}
	return RecentList
}

func (m *Storage) GetResList(ownerID int) []models.Resturant {
	sql := `Select id,name,address,img_link,description from resturants where owner_id=$1`

	rows, err := m.Db.Query(sql, ownerID)
	if err != nil {
		panic(err)
	}
	var sliceOfRes []models.Resturant
	for rows.Next() {
		var resDetails models.Resturant
		err := rows.Scan(&resDetails.Id, &resDetails.Name, &resDetails.Address, &resDetails.Img_link, &resDetails.Description)
		if err != nil {
			panic(err)
		}
		sliceOfRes = append(sliceOfRes, resDetails)
	}
	return sliceOfRes

}

func (m *Storage) DeleteResFromData(res_id int) {
	sql := `Delete from resturants where id=$1`

	_, err := m.Db.Exec(sql, res_id)
	if err != nil {
		panic(err)
	}
}
func (m *Storage) DeleteResOrders(res_id int) {
	sql := `Delete from ordered_items Where meal_id in ( select id from meals where res_id=$1)`
	_, err := m.Db.Exec(sql, res_id)
	if err != nil {
		panic(err)
	}

}
func (m *Storage) DeleteMeals(res_id int) {
	sql := `Delete from meals where res_id=$1`
	_, err := m.Db.Exec(sql, res_id)

	if err != nil {
		panic(err)
	}
}

func (m *Storage) CheckOwner(id int, res_id int) bool {
	sql := `select id from resturants where owner_id=$1`
	rows, err := m.Db.Query(sql, id)
	if err != nil {
		return false
	}
	for rows.Next() {
		var id int
		rows.Scan((&id))
		if id == res_id {
			return true
		}
	}
	return false

}

func (m *Storage) SingleResDetails(id int) models.UpdateResturant {
	sql := `select name,address,img_link,description from resturants where id=$1`
	row := m.Db.QueryRow(sql, id)
	var UpdateRes models.UpdateResturant
	err := row.Scan(&UpdateRes.Name, &UpdateRes.Address, &UpdateRes.Img_link, &UpdateRes.Description)
	if err != nil {
		panic(err)
	}
	return UpdateRes
}

func (m *Storage) UpdateRes(name string, address string, img_link string, desc string, res_id int) {
	sql := `Update resturants set name=$1,address=$2,img_link=$3,description=$4 where id=$5`
	_, err := m.Db.Exec(sql, name, address, img_link, desc, res_id)
	if err != nil {
		panic(err)
	}
}
