package main

import (
	"os"

	"github.com/ArnulfoVargas/echo_api/routes"
	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  godotenv.Load()

	e := echo.New()

  // Static files
  e.Static("/public", "public")

  //Middlewares
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
  }))

  // Routes
  routes.HandleRoutes(e)

  e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
