package router

import(
	"kafka/message"
	"service"
	"github.com/labstack/echo"
	"net/http"
)

func WriteLogs(c echo.Context) error{
	l := new(message.LogMessage)

	if err := c.Bind(l);  err !=nil{
		return c.JSON(http.StatusBadRequest,err)
	}

	err := service.Send(service.LogTopic, *l)
	if err !=nil{
		return c.JSON(http.StatusInternalServerError,err)
	}

	return c.JSON(http.StatusOK, l)
}

