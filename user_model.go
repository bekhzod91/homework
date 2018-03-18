package main

type Users struct {
	Id   int    `orm:"auto"`
	Email string `orm:"size(255)"`
	FirstName string `orm:"size(255)"`
	SecondName string `orm:"size(255)"`
}
