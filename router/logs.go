package router

import (
	"github.com/labstack/echo"
	"github.com/sukstar76/go-kafka/kafka/message"
	"github.com/sukstar76/go-kafka/service"
	"net/http"
)

func (h *Handler) WriteLogs(c echo.Context) error {
	l := new(message.LogMessage)

	if err := c.Bind(l); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.ls.Send(service.LogTopic, *l)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, l)
}
