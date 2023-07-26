package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) PostUser(name string, email string, username string, password string) {
	hasshedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	sql := `INSERT INTO users (name,email,username,password) values($1,$2,$3,$4)`
	_, err = s.Db.Exec(sql, name, email, username, hasshedpassword)
	if err != nil {
		panic(err)
	}
}
func (s *Storage) GetUserId(email string) int {
	sql := `SELECT id FROM users WHERE email=$1`
	var id int
	row := s.Db.QueryRow(sql, email)
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

func (s *Storage) ConfirmUser(email string, password string) (string, bool, string) {
	sql := `select password,'Customer' as user_type from users
	where email=$1
	union 
	select password,'Owner' as user_type from owner
	where email=$1`
	row := s.Db.QueryRow(sql, email)
	var hasshedpassword string
	var userType string
	err := row.Scan(&hasshedpassword, &userType)
	if err != nil {
		fmt.Println("line 41")
		return "username or password inncorrect", false, "not a valid person"
	}
	err = bcrypt.CompareHashAndPassword([]byte(hasshedpassword), []byte(password))
	if err != nil {
		fmt.Println("line 45")

		return "username or password inncorrect", false, "not a valid person"
	}

	return "loggedin", true, userType
}
func (s *Storage) ValidateUser(id float64) bool {
	sql := `SELECT username FROM users Where id=$1`
	var username string
	row := s.Db.QueryRow(sql, id)
	err := row.Scan(&username)
	if err != nil {
		return false
	}
	return true

}
