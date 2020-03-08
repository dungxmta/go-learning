package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	redis := driverRedis.Connect(redisAddr, redisPwd, redisDB)
	fmt.Println(redis)

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
	redis := driverRedis.GetInstance()

	c.Logger().Info("data:", data)

	result := redis.Client.LPush(QUEUE, data)
	if err := result.Err(); err != nil {
		c.Logger().Error(err)
	}
	return c.String(http.StatusOK, data)
}
