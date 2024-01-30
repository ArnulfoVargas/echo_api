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
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo/options"
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
  findOptions := options.Find()
  cursor, err := database.CategoryCollection.Find(context.TODO(), bson.D{}, findOptions.SetSort(bson.D{{"_id", -1}}))
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
      Message: "Error Fetching Data",
    })
  }

  return c.JSON(http.StatusOK, results)
}

func getCategoriesById(c echo.Context) error {
  objID, err := primitive.ObjectIDFromHex(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, dto.GenericDTO{
      Status: "Error",
      Message: "Bad Request",
    })
  }

  result := bson.M{}
  filter := bson.M{"_id" : objID}

  if err = database.CategoryCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Unexpected Error",
    })
  }

  return c.JSON(http.StatusOK, result)
}

func putCategories(c echo.Context) error {
  body := dto.Category{}
  decoder := json.NewDecoder(c.Request().Body)

  if err := decoder.Decode(&body); err != nil || len(body.Name) < 1 {
    return c.JSON(http.StatusBadRequest, dto.GenericDTO{
      Status: "Error",
      Message: "Bad Request",
    })
  }

  result := bson.M{}
  id := c.Param("id")
  objId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return c.JSON(http.StatusNotFound, dto.GenericDTO{
      Status: "Error",
      Message: "Not found",
    })
  }

  if err := database.
  CategoryCollection.
  FindOne(context.TODO(), bson.M{"_id" : objId}).
  Decode(&result); err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Unexpected Error",
    })
  }

  register := map[string]any {
    "name" : body.Name,
    "slug" : utils.CreateSlug(body.Name),
  }
  updateString := bson.M{
    "$set" : register,
  }
  _, err = database.CategoryCollection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": objId}}, updateString)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Unexpected Error",
    })
  }

  return c.JSON(http.StatusOK, dto.GenericDTO{
    Status: "Ok",
    Message: "Successfully Updated",
  })
}

func deleteCategory(c echo.Context) error {
  body := dto.Category{}
  decoder := json.NewDecoder(c.Request().Body)

  if err := decoder.Decode(&body); err != nil || len(body.Name) < 1 {
    return c.JSON(http.StatusBadRequest, dto.GenericDTO{
      Status: "Error",
      Message: "Bad Request",
    })
  }

  result := bson.M{}
  id := c.Param("id")
  objId, err := primitive.ObjectIDFromHex(id)
  if err != nil {
    return c.JSON(http.StatusNotFound, dto.GenericDTO{
      Status: "Error",
      Message: "Not found",
    })
  }

  if err := database.
  CategoryCollection.
  FindOne(context.TODO(), bson.M{"_id" : objId}).
  Decode(&result); err != nil {
    return c.JSON(http.StatusInternalServerError, dto.GenericDTO{
      Status: "Error",
      Message: "Unexpected Error",
    })
  }

  database.CategoryCollection.DeleteOne(context.TODO(), bson.M{"_id":objId})

  return c.JSON(http.StatusOK, dto.GenericDTO{
    Status: "Ok",
    Message: "Deleted Successfully",
  })
}
