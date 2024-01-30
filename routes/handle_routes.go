package routes

import (
	categories_routes "github.com/ArnulfoVargas/echo_api/routes/categories"
	"github.com/ArnulfoVargas/echo_api/routes/example"
	products_routes "github.com/ArnulfoVargas/echo_api/routes/products"
	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo) {
  example.HandleExample(e)
  categories_routes.HanderCategories(e)
  products_routes.HandleProducts(e)
}
