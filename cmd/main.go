package main

import (
	//"fmt"
	"github.com/bpietrzakk/swift_codes/internal/api"
	"github.com/bpietrzakk/swift_codes/internal/database"
	"github.com/bpietrzakk/swift_codes/internal/parser"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect() // connect to database and create tables in database

	// parse the data from csv data
	parserData , err := parser.ParseSwiftCodesCSV("internal/data/Interns_2025_SWIFT_CODES - Sheet1.csv")
	if err != nil {
		panic(err)
	}

	database.LoadSwiftCodesToDB(parserData)


	// --------------API--------------
	// start api
	r := gin.Default()

	// add endopints

	api.RegisterEndpoint1(r) // endpoint 1
	api.RegisterEndpoint2(r) // endpoint 2

	r.GET("/hello", func(c *gin.Context) {
		// var example []models.SwiftCode
		// database.DB.Limit(1).Find(&example)
		// c.JSON(200, example)
		c.JSON(200, gin.H{
			"message":"Witaj świecie!",
		})
	})

	r.Run(":8080")
}