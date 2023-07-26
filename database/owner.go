package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *Storage) ConfirmOwner(email string, password string) (string, bool) {
	sql := `Select password from owner where email=$1`
	row := s.Db.QueryRow(sql, email)
	var hashedpassword string
	err := row.Scan(&hashedpassword)
	if err != nil {
		return "username or password incorrect", false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		return "username or password incorrect", false
	}
	return "loggedin", true
}

func (s *Storage) GetOwnerId(email string) int {
	sql := `select id from owner where email=$1`
	var id int
	row := s.Db.QueryRow(sql, email)
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func (s *Storage) CreateOwner(name string, email string, username string, password string) {
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	sql := `Insert Into owner (name,email,username,password) values ($1,$2,$3,$4)`
	_, err := s.Db.Exec(sql, name, email, username, hashedpassword)
	if err != nil {
		panic(err)
	}

}

func (s *Storage) Validateowner(id int) bool {
	sql := `select username from owner where id=$1`
	_, err := s.Db.Exec(sql, id)
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) UpdateOwner(name string, email string, usename string, password string, ownerId int) {
	sql := `Update owner Set name =$1,email=$2,username=$3,password=$4 where id=$5`
	hassedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	fmt.Println(ownerId, "kcekn")
	s.Db.Exec(sql, name, email, usename, hassedPassword, ownerId)
}
