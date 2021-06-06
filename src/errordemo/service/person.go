package service

import (
	"errordemo/dao"
	"errordemo/po"
 	"github.com/pkg/errors"
	"database/sql"
	// "fmt"
)

func SelectPerson(userId int) ([]po.Person, error) {
	person, err := dao.SelectPerson(userId)
	if err != nil {
		return nil, err
	}
	return person, nil
}


func QueryPerson(userId int) (string, error){
	userName, err := dao.QueryPerson(userId)
	rootErr := errors.Cause(err)
	
	//此处可以根据业务需求，对errNoRows做判断，返回空字符串，或者err
	if errors.Is(rootErr, sql.ErrNoRows) {
		return "", nil
	}
	return userName, err
}