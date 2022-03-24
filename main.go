package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var e *echo.Echo

const envFilename = ".env"

var config *Config

func init() {
	config = new(Config)

	config.Configurar()
}

func main() {
	e = echo.New()

	go RodarRotina()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.POST},
	}))

	e.POST("/mensageria", func(c echo.Context) error {
		args := new(ArgumentosEmail)

		if err := c.Bind(args); err != nil {
			return c.JSON(http.StatusNotAcceptable, map[string]interface{}{"msg": "JSON inválido"})
		}

		if err := args.ValidarDados(); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": err.Error()})
		}

		if err := args.AdicionarFila(); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": fmt.Sprintf("Falha ao adicionar a fila: %s", err.Error())})
		}

		return c.JSON(http.StatusOK, "OK")
	})

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}
