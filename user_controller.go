package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {}

type RequestUserData struct {
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
}

func (ctrl UserController) list (c *gin.Context) {
	o := orm.NewOrm()
	qs := o.QueryTable("users")

	// List
	users := []Users{}
	qs.All(&users, "id", "email")

	// Count
	count, _ := qs.Count()

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"data": &users,
	})
}

func (ctrl UserController) detail (c *gin.Context) {
	id := c.Param("id")
	idInt, errInt := strconv.Atoi(id)

	if errInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Incorrect type. Expected pk value, received " + id + ".",
		})

		return
	}

	user := Users{Id: idInt }
	o := orm.NewOrm()
	err := o.Read(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "Invalid pk " + id + " - object does not exist.",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{ "data": &user, })

}

func (ctrl UserController) create (c *gin.Context) {
	data := new(RequestUserData)
	c.ShouldBindWith(&data, binding.JSON)

	isValid, errors := UserForm(data)

	if isValid {
		user := new(Users)
		user.Email = data.Email
		user.FirstName = data.FirstName
		user.SecondName = data.SecondName

		o := orm.NewOrm()
		o.Insert(user)

		c.JSON(http.StatusCreated, gin.H {
			"data": user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors": errors,
		})
	}
}

func (ctrl UserController) update (c *gin.Context) {
	id := c.Param("id")
	idInt, errInt := strconv.Atoi(id)

	if errInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Incorrect type. Expected pk value, received " + id + ".",
		})

		return
	}

	user := Users{Id: idInt }
	o := orm.NewOrm()
	err := o.Read(&user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "Invalid pk " + id + " - object does not exist.",
		})

		return
	}

	data := new(RequestUserData)
	c.ShouldBindWith(&data, binding.JSON)

	isValid, errors := UserForm(data)

	if isValid {
		user.Email = data.Email
		user.FirstName = data.FirstName
		user.SecondName = data.SecondName

		o.Update(&user)

		c.JSON(http.StatusOK, gin.H{
			"data": &user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H {
			"errors": errors,
		})
	}
}

func (ctrl UserController) delete (c *gin.Context) {
	id := c.Param("id")
	idInt, errInt := strconv.Atoi(id)

	if errInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "Incorrect type. Expected pk value, received " + id + ".",
		})

		return
	}

	user := Users{Id: idInt }

	o := orm.NewOrm()
	o.Delete(&user)

	c.JSON(http.StatusNoContent, nil)
}