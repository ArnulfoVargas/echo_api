package categories_routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ArnulfoVargas/echo_api/database"
	"github.com/ArnulfoVargas/echo_api/models/dto"
	"github.com/ArnulfoVargas/echo_api/utils"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func postCategory(c echo.Context) error {
  var bodyDTO dto.Category

  if err := json.NewDecoder(c.Request().Body).Decode(&bodyDTO); 
    err != nil || len(bodyDTO.Name) == 0{
    return c.JSON(http.StatusBadRequest, dto.GenericDTO{
      Status: "Error",
      Message: "Bad Request",
    })
  }

  register := bson.D{ {"name", bodyDTO.Name}, {"slug", utils.CreateSlug(bodyDTO.Name)} }
  _, err := database.CategoryCollection.InsertOne(context.TODO(), register)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Error Storing Data",
    })
  }

  return c.JSON(http.StatusCreated, dto.GenericDTO{
    Status: "Ok",
    Message: "Stored Successfully", 
  })
}

func getCategories(c echo.Context) error {

  return nil
}
