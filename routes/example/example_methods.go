package example

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	dto_example "github.com/ArnulfoVargas/echo_api/models/dto"
	"github.com/labstack/echo/v4"
)

func getExample(c echo.Context) error {
  return c.JSON(http.StatusOK, echo.Map{
    "status"  : http.StatusOK,
    "message" : "Sended from Echo framework",
    "header"  : c.Request().Header.Get("Authorization"),
  })
}

func postExample(c echo.Context) error {
  categoryDto := dto_example.Category{}
  decoder := json.NewDecoder(c.Request().Body)

  if err := decoder.Decode(&categoryDto); err != nil{
    return err
  }

  return c.JSON(http.StatusOK, echo.Map{
    "status" : http.StatusOK,
    "message": "Method Post",
    "name"   : categoryDto.Name,
  })
}

func putExample(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, "Put from echo | id = " + id)
}

func deleteExample(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, "Delete from echo | id " + id)
}

func getParamsExample(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, "Get id = " + id)
}

func getQueryExample(c echo.Context) error { 
  id := c.QueryParam("id")
  return c.String(http.StatusOK, "Get query | id = " + id)
}

func uploadFileExample(c echo.Context) error {
  file, err  := c.FormFile("image")
  if err != nil {
    return err
  }

  src, err := file.Open()
  if err != nil {
    return err
  }
  defer src.Close()

  // Give  the file a unique name
  splitFileName := strings.Split(file.Filename, ".")
  extension := splitFileName[len(splitFileName) - 1] 
  currentDate := strings.Split(time.Now().String(), " ")
  fileName := string(currentDate[4][6:14]) + "." + extension
  fileDir := "public/uploads/images/" + fileName

  dst, err := os.Create(fileDir)
  if err != nil {
    return err
  }
  defer dst.Close()

  _, err = io.Copy(dst, src)
  if err != nil {
    return err
  }

  return c.JSON(http.StatusCreated, echo.Map{
    "status" : http.StatusCreated,
    "message": "file successfully uploaded",
  })
}
