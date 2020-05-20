package router

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sukstar76/go-kafka/model"
	"net/http"
)

func (h *Handler) SaveUser(c echo.Context) error {
	ctx := c.Request().Context()

	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	user, err := h.us.Create(&ctx, u)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single documents : ", user)
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	user, err := h.us.GetID(&ctx, id)

	if err != nil {
		return c.String(http.StatusNotFound, id+" not exits")
	}

	return c.JSON(http.StatusOK, user)

}

func (h *Handler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if _, err := h.us.GetID(&ctx, id); err != nil {
		return c.String(http.StatusNotFound, id+" not exist")

	}
	user, err := h.us.Update(&ctx, u)

	if err != nil {
		return err
	}

	fmt.Println("Updated user : ", user)
	return c.JSON(http.StatusOK, u)
}
