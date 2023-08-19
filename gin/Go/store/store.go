package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rzeAkbari/GoGo/gin/server"
)

type Store struct {
	DB *sql.DB
}

func NewStore(connection string) (*Store, error) {
	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		return nil, err
	}
	return &Store{
		db,
	}, nil
}

func (s *Store) CheckAuth(credentials server.Credentials) (bool, error) {
	stmt, err := s.DB.Prepare("select id from user where username = ? AND password=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var id string
	err = stmt.QueryRow(credentials.UserName, credentials.Password).Scan(&id)
	if err != nil {
		return false, err
	}
	if id != "" {
		return true, nil
	}

	return false, nil
}
