package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"testProject/learning/example/api_jwt_mongo/config"
	// driverMongo "testProject/learning/example/api_jwt_mongo/driver/mongo"
	driverRedis "testProject/learning/example/api_jwt_mongo/driver/redis"
)

const (
	QUEUE = "plugin_config"
)

func main() {

	redisAddr := config.GetStr("REDIS_ADDR")
	redisPwd := config.GetStr("REDIS_PWD")
	redisDB := int(config.Get("REDIS_DB").(float64))

	fmt.Println(redisAddr)

	msgQueue, err := driverRedis.GetInstance().Init(redisAddr, redisPwd, redisDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msgQueue)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", test)
	e.GET("/push/:data", push)

	e.Logger.Fatal(e.Start(config.GetStr("REST_APP")))
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func push(c echo.Context) error {
	data := c.Param("data")

	c.Logger().Info("data:", data)

	_, err := driverRedis.GetInstance().LPush(QUEUE, data)
	// if err := result.Err(); err != nil {
	if err != nil {
		c.Logger().Error(err)
	}
	return c.String(http.StatusOK, data)
}
