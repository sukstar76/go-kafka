package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func UserHandler(v1 *echo.Group){
	user := v1.Group("/users")
	user.GET("",func(c echo.Context) error{
		return c.String(http.StatusOK,"User Page\n")
	})
	user.POST("",SaveUser)
	user.GET("/:id",GetUser)
	user.PUT("/:id",UpdateUser)
//	user.DELETE("/:id",h.DeleteUser)

	log := v1.Group("/logs")
	log.POST("",WriteLogs)
}
