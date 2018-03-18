package main

import (
	"github.com/astaxie/beego/orm"
	"testing"
)

type ResponseError struct {
	Detail string `json:"detail"`
}

type UserDataResponse struct {
	Id int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
}

type UserListResponse struct {
	Count int `json:"count"`
	Data []UserDataResponse `json:"data"`
}

type UserDetailResponse struct {
	Data UserDataResponse `json:"data"`
}

type CreateOrUpdateError struct {
	Errors struct {
		Email string `json:"email"`
		FirstName string `json:"firstName"`
		SecondName string `json:"secondName"`
	} `json:"errors"`
}


func initDB(dbName string) {
	// register model
	orm.RegisterModel(new(Users))
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	// set default database
	err:=orm.RegisterDataBase("default", "sqlite3", "file:" + dbName)
	if err!=nil{
		panic(err)
	}

	// create table
	orm.RunSyncdb("default", false, true)
}

func setupSubTest(t *testing.T, handler func()) func(t *testing.T) {
	o := orm.NewOrm()
	o.Begin()
	handler()
	return func(t *testing.T) {
		o.Rollback()
	}
}
