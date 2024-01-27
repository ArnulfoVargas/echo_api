package categories_routes

import "github.com/labstack/echo/v4"

func HanderCategories(e *echo.Echo) {
  prefix := "/api/categories"
  e.POST(prefix, postCategory)
  e.GET(prefix, getCategories)
}
