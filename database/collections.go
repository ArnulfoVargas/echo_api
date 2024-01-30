package database

var CategoryCollection = MongoClient.Database(DbName).Collection("categories")
var ProductCollection = MongoClient.Database(DbName).Collection("products")
