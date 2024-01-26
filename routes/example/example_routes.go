package example

import (
	echo "github.com/labstack/echo/v4"
)

const prefix string = "/api/example"

func HandleExample(e *echo.Echo) {
  e.GET(prefix, getExample) 
  e.GET(prefix + "/:id", getParamsExample)
  e.POST(prefix, postExample)
  e.PUT(prefix + "/:id", putExample)
  e.DELETE(prefix + "/:id", deleteExample)
  e.POST(prefix + "/upload", uploadFileExample)

  e.GET(prefix + "/query", getQueryExample)
}
