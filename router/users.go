package router

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/sukstar76/model"
	"github.com/sukstar76/service"
)


func SaveUser(c echo.Context) error{
	ctx := c.Request().Context()

	u:= new(model.User)

	if err := c.Bind(u); err!=nil{
		return err
	}

	user, err := service.Create(&ctx,u)
	if err != nil{
		return err
	}

	fmt.Println("Inserted a single documents : " , user)
	return c.JSON(http.StatusCreated,u)
}

func GetUser(c echo.Context) error{
	ctx := c.Request().Context()

	id := c.Param("id")

	user,err:=service.GetID(&ctx,id)

	if err!=nil {
		return c.String(http.StatusNotFound, id+ " not exits")
	}

	return c.JSON(http.StatusOK,user)

}

func UpdateUser(c echo.Context) error{
	ctx := c.Request().Context()

	id := c.Param("id")

	u := new(model.User)

	if err := c.Bind(u); err !=nil{
		return err
	}

	if _, err := service.GetID(&ctx, id); err != nil{
		return c.String(http.StatusNotFound,id+" not exist")

	}
	user, err := service.Update(&ctx, u)

	if err != nil{
		return err
	}

	fmt.Println("Updated user : ", user)
	return c.JSON(http.StatusOK, u)
}


