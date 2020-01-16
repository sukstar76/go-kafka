package main

import(
	"net/http"
	"github.com/sukstar76/go-kafka/db"
	"github.com/sukstar76/go-kafka/router"
	"github.com/sukstar76/go-kafka/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sukstar76/go-kafka/kafka"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	e :=echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	client, err := db.Connect()

	if err !=nil{
		e.Logger.Fatal(err)
	}

	service.DBCollection = client.Database("test1").Collection("user")
	kafka.SetTopic()

	producer := kafka.ConnectProducer()

	service.LogService = producer

	e.GET("/",func(c echo.Context) error{
		return c.String(http.StatusOK,"Hello\n")
	})
	v1 := e.Group("/api")
	router.UserHandler(v1)

	consumer, err := kafka.ConsumerConnect()
	if err!=nil{
		e.Logger.Fatal(err)
	}

	go consumer.Consume()
	
	e.Logger.Fatal(e.Start(":8080"))

}

