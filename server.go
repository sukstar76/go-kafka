package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sukstar76/go-kafka/db"
	"github.com/sukstar76/go-kafka/kafka"
	"github.com/sukstar76/go-kafka/router"
	"github.com/sukstar76/go-kafka/service"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client, err := db.Connect()

	if err != nil {
		e.Logger.Fatal(err)
	}

	us := service.NewUserService(client)

	kafka.SetTopic()

	producer := kafka.ConnectProducer()

	ls := service.NewLogService(producer)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello\n")
	})
	v1 := e.Group("/api")
	h := router.NewHandler(us, ls)

	h.Register(v1)

	consumer, err := kafka.ConsumerConnect()
	if err != nil {
		e.Logger.Fatal(err)
	}

	go consumer.Consume()

	e.Logger.Fatal(e.Start(":8080"))

}
