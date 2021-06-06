package dao

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"errordemo/po"
	"github.com/pkg/errors"
)

var Db *sqlx.DB

func init() {
	username := "root"
	ip := "127.0.0.1"
	password := "123456"
	port := 3306
	database := "test"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, ip, port, database)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open Mysql error", err)
		return 
	}
	Db = db
}

func SelectPerson(userId int) ([]po.Person, error){
	var person []po.Person
	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", userId)
	if err != nil {
		fmt.Println("select error:", err)
		return person, err
	}
	fmt.Println("select success:", person)
	return person, nil
}


//作业
func QueryPerson(userId int) (string, error){
	var userName string
	err := Db.QueryRow("select user_id, username, sex, email from person where user_id=?", userId).Scan(&userName)

	if err != nil {
		return "", errors.Wrap(err, "ErrNoRows")
	}
	return userName, nil
}

