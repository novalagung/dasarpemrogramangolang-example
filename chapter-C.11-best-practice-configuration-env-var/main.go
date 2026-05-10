package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	confAppName := os.Getenv("APP_NAME")
	if confAppName == "" {
		e.Logger.Fatal("APP_NAME config is required")
	}

	confServerPort := os.Getenv("SERVER_PORT")
	if confServerPort == "" {
		e.Logger.Fatal("SERVER_PORT config is required")
	}

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	server := new(http.Server)
	server.Addr = ":" + confServerPort

	if confServerReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); confServerReadTimeout != "" {
		duration, _ := strconv.Atoi(confServerReadTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}

	if confServerWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); confServerWriteTimeout != "" {
		duration, _ := strconv.Atoi(confServerWriteTimeout)
		server.WriteTimeout = time.Duration(duration) * time.Minute
	}

	e.Logger.Print("Starting", confAppName)
	e.Logger.Fatal(e.StartServer(server))
}
