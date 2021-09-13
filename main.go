package main

import (
	"fmt"

	"net/http"
	"github.com/labstack/echo"
	redis "github.com/go-redis/redis"
)

func initializeRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 접근 url 및 port
		Password: "",               // password ""값은 없다는 뜻
		DB:       0,                // 기본 DB 사용
	})

	_, err := client.Ping().Result()
	return client, err
}

var connection_client *redis.Client
var connection_err error

func main() {
	connection_client, connection_err = initializeRedisClient()
	e := echo.New()
	e.GET("/upload/:time", getTime)
	e.POST("/upload/:time", setTime)
	e.DELETE("/upload/:time", deleteTime)
	e.Logger.Fatal(e.Start(":1331"))
}

func setTime(c echo.Context) error {
	time := c.Param("time")
	data := c.FormValue("data")
	fmt.Println(time, data, connection_err, connection_client)
	if connection_err == nil {
		connection_client.Set(time, data, 0)
		return c.JSON(http.StatusOK, true)
	}else {
		return c.JSON(http.StatusNotFound, false)
	}
}
func getTime(c echo.Context) error {
	time := c.Param("time")
	data, err := connection_client.Get(time).Result()
	fmt.Println(time, data, err)
	return c.JSON(http.StatusOK, data)
}


func deleteTime(c echo.Context) error {
	time := c.Param("time")
	res, err := connection_client.Do("del", time).Result()
	if connection_err == nil {
		if res == 0 {
			return c.JSON(http.StatusOK, "해당 시간에 로그정보가 존재하지 않습니다.")
		}else{
			return c.JSON(http.StatusOK, "해당시간에 로그가 삭제되었습니다.")
		}
	}else {
		return c.JSON(http.StatusNotFound, err)
	}
}