package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Register(v1 *echo.Group) {
	user := v1.Group("/users")
	user.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "User Page\n")
	})
	user.POST("", h.SaveUser)
	user.GET("/:id", h.GetUser)
	user.PUT("/:id", h.UpdateUser)
	//	user.DELETE("/:id",h.DeleteUser)

	log := v1.Group("/logs")
	log.POST("", h.WriteLogs)
}
