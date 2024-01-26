package routes

import (
	"github.com/ArnulfoVargas/echo_api/routes/example"
	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {
  example.HandleExample(e)
}
