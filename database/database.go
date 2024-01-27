package database

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = InitDatabase()
var clientOpts = options.Client().ApplyURI(generateURI())
var DbName = ""

func InitDatabase() *mongo.Client {
  ctx := context.TODO()
  client, err := mongo.Connect(ctx, clientOpts)
  if err != nil {
    panic("Error while connecting to the database")
  }

  err = client.Ping(ctx, nil)
  if err != nil {
    panic("Error while connecting to the database")
  }

  return client
}

func IsConnected() bool {
  err := MongoClient.Ping(context.TODO(), nil)
  if err != nil {
    return false
  }

  return true
}

func generateURI() string {
  err := godotenv.Load()
  DbName = os.Getenv("DB_NAME")

  if err != nil {
    panic("Error loading .env")
  }

  sb := strings.Builder{}

  sb.WriteString("mongodb://")
  sb.WriteString(os.Getenv("DB_HOST"))
  sb.WriteByte(':')
  sb.WriteString(os.Getenv("DB_PORT"))
  sb.WriteByte('/')
  sb.WriteString(DbName)

  uri := sb.String()
  fmt.Println("Connected at: " + uri)

  return uri
}
