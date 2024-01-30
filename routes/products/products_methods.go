package products_routes

import (
  "context"
  "encoding/json"
  "net/http"

  "github.com/ArnulfoVargas/echo_api/database"
  "github.com/ArnulfoVargas/echo_api/models/dto"
  "github.com/ArnulfoVargas/echo_api/utils"
  "github.com/labstack/echo/v4"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func postProducts(c echo.Context) error {
  var bodyDTO dto.ProductDTO

  if err := json.NewDecoder(c.Request().Body).Decode(&bodyDTO); 
  err != nil || len(bodyDTO.Name) == 0{
    return c.JSON(http.StatusBadRequest, dto.GenericDTO{
      Status: "Error",
      Message: "Bad Request",
    })
  }

  categoryId, err := primitive.ObjectIDFromHex(bodyDTO.CategoryID)
  if err != nil {
    return err
  }

  register := bson.D{ 
    {"name", bodyDTO.Name}, 
    {"slug", utils.CreateSlug(bodyDTO.Name)},
    {"stock", bodyDTO.Stock},
    {"price", bodyDTO.Price},
    {"description", bodyDTO.Description},
    {"category_id", categoryId},
  }
  _, err = database.ProductCollection.InsertOne(context.TODO(), register)
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

func getProducts(c echo.Context) error {
  pipeline := []bson.M{
    bson.M{"$match" : bson.M{}},
    bson.M{"$lookup": bson.M{"from": "categories", "localField": "category_id", "foreignField":"_id", "as":"category"}},
    bson.M{"$sort":bson.M{"_id": -1}},
  }

  cursor, err := database.ProductCollection.Aggregate(context.TODO(), pipeline)
  if err != nil{
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Error Fetching Data",
    })
  }

  results := []bson.M{}
  if err = cursor.All(context.TODO(), &results); err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Error Parsing Data",
    })
  }

  return c.JSON(http.StatusOK, results)
}
