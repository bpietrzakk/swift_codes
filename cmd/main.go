package main

import (
	//"fmt"
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/bpietrzakk/swift_codes/internal/models"
	"github.com/bpietrzakk/swift_codes/internal/database"
	"github.com/bpietrzakk/swift_codes/internal/models"
	"github.com/bpietrzakk/swift_codes/internal/parser"
)

func main() {
	database.Connect() // connect to database and create tables in database

	// parse the data from csv data
	parserData , err := parser.ParseSwiftCodesCSV("../internal/data/test.csv")
	if err != nil {
		panic(err)
	}

	for _, record := range parserData {
		database.DB.Create(&record)
	}
	fmt.Println("data has loaded to database! ")




	// start api
	r := gin.Default()

	// simple endpoint 
	r.GET("/", func(c *gin.Context) {
		var example []models.SwiftCode
		database.DB.Limit(1).Find(&example)
		c.JSON(200, example)
	})

	r.Run(":8080")
}