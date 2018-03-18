package main

import (
	"github.com/astaxie/beego/orm"
	"testing"
)

type ResponseError struct {
	Detail string `json:"Detail"`
}

type UserDataResponse struct {
	Id int `json:"Id"`
	Email string `json:"Email"`
	FirstName string `json:"FirstName"`
	SecondName string `json:"SecondName"`
}

type UserListResponse struct {
	Count int `json:"Count"`
	Data []UserDataResponse `json:"data"`
}

type UserDetailResponse struct {
	Data UserDataResponse `json:"data"`
}

type CreateOrUpdateError struct {
	Errors struct {
		Email string `json:"Email"`
		FirstName string `json:"FirstName"`
		SecondName string `json:"SecondName"`
	} `json:"Errors"`
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
