package products_routes

import "github.com/labstack/echo/v4"

func HandleProducts(r *echo.Echo) {
  prefix := "/api/products"
  r.POST(prefix, postProducts)
  r.GET(prefix, getProducts)
}
