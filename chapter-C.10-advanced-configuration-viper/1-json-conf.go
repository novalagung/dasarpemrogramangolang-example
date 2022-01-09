package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()

	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Print("Starting", viper.GetString("appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}
