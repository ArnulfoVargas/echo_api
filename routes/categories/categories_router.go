package categories_routes

import "github.com/labstack/echo/v4"

func HanderCategories(e *echo.Echo) {
  prefix := "/api/categories"
  e.POST(prefix, postCategory)
  e.GET(prefix, getCategories)
  e.GET(prefix + "/:id", getCategoriesById)
  e.PUT(prefix + "/:id", putCategories)
  e.DELETE(prefix + "/:id", deleteCategory)
}
