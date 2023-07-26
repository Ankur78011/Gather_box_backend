package handlers

import (
	"database/sql"

	"gatherbox.com/database"
)

type ApiHandler struct {
	Db      *sql.DB
	Storage *database.Storage
}

func NewApiHandler(db *sql.DB) *ApiHandler {
	return &ApiHandler{Db: db, Storage: database.NewStorage(db)}
}
