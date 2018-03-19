package main

import (
	"os"
	"bytes"
	"testing"
	"net/http"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"

)

var route = setupRouter()

func TestMain(m *testing.M) {
	dbname := "test.db"

	// Clear test database
	os.Remove(dbname)

	// Init new database
	initDB(dbname)

	retCode := m.Run()
	os.Exit(retCode)
}

func initData() {
	o := orm.NewOrm()
	users := []Users{
		{ Id: 1, Email: "admin@example.com", FirstName: "FirstName", SecondName: "SecondName" },
		{ Id: 2, Email: "moder@example.com", FirstName: "ModerFirstName", SecondName: "ModerSecondName" },
	}

	o.InsertMulti(2, &users)
}

func initTest(t *testing.T) Client {
	teardownTestCase := setupSubTest(t, initData)
	defer teardownTestCase(t)

	client := Client{ route }

	return client
}


func TestUserList(t *testing.T) {
	client := initTest(t)

	response := client.get("/api/v1/user/")

	responseData := UserListResponse{}
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusOK)

	// First row
	assert.Equal(t, responseData.Data[0].Id, 1)
	assert.Equal(t, responseData.Data[0].Email, "admin@example.com")

	// Second row
	assert.Equal(t, responseData.Data[1].Id, 2)
	assert.Equal(t, responseData.Data[1].Email, "moder@example.com")
}

func TestUserListEmpty(t *testing.T) {
	client := initTest(t)

	o := orm.NewOrm()
	qs := o.QueryTable("users")
	qs.Filter("id__gt", 0).Delete()

	response := client.get("/api/v1/user/")

	responseData := UserListResponse{}
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, responseData.Count, 0)
	assert.Len(t, responseData.Data, 0)
}

func TestUserDetail(t *testing.T) {
	client := initTest(t)

	response := client.get("/api/v1/user/1/")

	responseData := new(UserDetailResponse)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, responseData.Data.Id, 1)
	assert.Equal(t, responseData.Data.Email, "admin@example.com")
	assert.Equal(t, responseData.Data.FirstName, "FirstName")
	assert.Equal(t, responseData.Data.SecondName, "SecondName")
}

func TestUserDetailNotFound(t *testing.T) {
	client := initTest(t)

	response := client.get("/api/v1/user/3/")
	assert.Equal(t, response.Code, http.StatusNotFound)
}

func TestUserCreate(t *testing.T)  {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{
		Email:      "user@example.com",
		FirstName:  "UserFirstName",
		SecondName: "UserSecondName",
	})
	body := bytes.NewBuffer(request)

	response := client.post("/api/v1/user/", body)

	responseData := new(UserDetailResponse)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	o := orm.NewOrm()
	user := Users{ Id: responseData.Data.Id }
	o.Read(&user)

	assert.Equal(t, response.Code, http.StatusCreated)
	assert.Equal(t, user.Email, "user@example.com")
	assert.Equal(t, user.FirstName, "UserFirstName")
	assert.Equal(t, user.SecondName, "UserSecondName")
}

func TestUserCreateEmptyError(t *testing.T)  {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{
		Email:      "",
		FirstName:  "",
		SecondName: "",
	})
	body := bytes.NewBuffer(request)

	response := client.post("/api/v1/user/", body)

	responseData := new(CreateOrUpdateError)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusBadRequest)
	assert.Equal(t, responseData.Errors.Email, "Enter a valid email address.")
	assert.Equal(t, responseData.Errors.FirstName, "This field may not be blank.")
	assert.Equal(t, responseData.Errors.SecondName, "This field may not be blank.")
}

func TestUserUpdateEmptyError(t *testing.T)  {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{
		Email:      "admin1example.com",
		FirstName:  "",
		SecondName: "",
	})
	body := bytes.NewBuffer(request)

	response := client.put("/api/v1/user/1/", body)

	responseData := new(CreateOrUpdateError)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusBadRequest)
	assert.Equal(t, responseData.Errors.Email, "Enter a valid email address.")
	assert.Equal(t, responseData.Errors.FirstName, "This field may not be blank.")
	assert.Equal(t, responseData.Errors.SecondName, "This field may not be blank.")
}


func TestUserUpdate(t *testing.T)  {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{
		Email:      "admin1@example.com",
		FirstName:  "FirstName1",
		SecondName: "SecondName2",
	})
	body := bytes.NewBuffer(request)

	response := client.put("/api/v1/user/1/", body)

	o := orm.NewOrm()
	user := Users{ Id: 1 }
	o.Read(&user)

	assert.Equal(t, response.Code, http.StatusOK)
	assert.Equal(t, user.Email, "admin1@example.com")
	assert.Equal(t, user.FirstName, "FirstName1")
	assert.Equal(t, user.SecondName, "SecondName2")
}

func TestUserUpdateNotFound(t *testing.T) {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{})
	body := bytes.NewBuffer(request)

	response := client.put("/api/v1/user/10/", body)

	responseData := new(ResponseError)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusNotFound)
	assert.Equal(t, responseData.Detail, "Invalid pk 10 - object does not exist.")
}

func TestUserUpdateWrongUrl(t *testing.T) {
	client := initTest(t)

	request, _ := json.Marshal(&UserDataResponse{})
	body := bytes.NewBuffer(request)

	response := client.put("/api/v1/user/a10/", body)

	responseData := new(ResponseError)
	json.Unmarshal([]byte(response.Body.String()), &responseData)

	assert.Equal(t, response.Code, http.StatusBadRequest)
	assert.Equal(t, responseData.Detail, "Incorrect type. Expected pk value, received a10.")
}

func TestUserDelete(t *testing.T) {
	client := initTest(t)

	response := client.delete("/api/v1/user/1/")
	assert.Equal(t, response.Code, http.StatusNoContent)

	o := orm.NewOrm()
	user := Users{ Id: 1 }
	err := o.Read(&user)

	assert.NotNil(t, err)
}