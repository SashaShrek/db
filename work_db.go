package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host   *string
	port   *int
	user   *string
	passwd *string
	dbname *string
)

/*Заполнить данные для авторизации в БД*/
func FillData(Host *string, Port *int, User *string, Password *string, DBname *string) {
	host = Host
	port = Port
	user = User
	passwd = Password
	dbname = DBname
}

/*Метод для выборки значений*/
func Select(query string) (*sql.Rows, *sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host='%s' port=%d user='%s' password='%s' dbname='%s' sslmode=disable",
		*host, *port, *user, *passwd, *dbname))
	if err != nil {
		return nil, nil, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	return rows, db, nil
}

/*Метод для выполнения insert, update, delete*/
func InsertOrUpdate(query string) error {
	db, err := sql.Open("postgres", fmt.Sprintf("host='%s' port=%d user='%s' password='%s' dbname='%s' sslmode=disable",
		*host, *port, *user, *passwd, *dbname))
	if err != nil {
		return err
	}
	defer db.Close()

	insert, err := db.Query(query)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}
