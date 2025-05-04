package api

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/bpietrzakk/swift_codes/internal/database"
	"github.com/bpietrzakk/swift_codes/internal/models"
	"github.com/bpietrzakk/swift_codes/internal/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ENDPOINT 1: GET: /v1/swift-codes/{swift-code}
//
//	response for hq with branches
//	response for branch single
func RegisterEndpoint1(r *gin.Engine) {
	r.GET("/v1/swift-codes/:swift_code", func(c *gin.Context) {
		swift_code := c.Param("swift_code")

		// check if swift code has 11 characters
		if utf8.RuneCountInString(swift_code) != 11 {
			c.JSON(400, gin.H{"error": "swift code must have 11 characters"})
			return
		}

		// find swift code in database
		var wanted_swift models.SwiftCode
		result := database.DB.Where("swift_code = ?", swift_code).First(&wanted_swift)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": "Swift code was not found"})
				return
			}
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}

		// if branch
		if !wanted_swift.IsHeadquarter {
			c.JSON(200, gin.H{
				"address":       wanted_swift.Address,
				"bankName":      wanted_swift.BankName,
				"countryISO2":   wanted_swift.CountryISO2,
				"countryName":   wanted_swift.CountryName,
				"isHeadquarter": wanted_swift.IsHeadquarter,
				"swiftCode":     wanted_swift.SwiftCode,
			})
		} else {
			// if headquarter
			prefix := wanted_swift.SwiftCode[:8]
			var branches []models.SwiftCode
			database.DB.Where("swift_code LIKE ? AND swift_code != ?", prefix+"%", wanted_swift.SwiftCode).Find(&branches)

			response := responses.Build_HQB_response(wanted_swift, branches)
			c.JSON(200, response)

		}

	})
}

func RegisterEndpoint2(r *gin.Engine) {
	r.GET("/v1/swift-codes/country/:countryISO2code", func(c *gin.Context) {
		countryISO2 := c.Param("countryISO2code")

		// check if iso2 has 2 charakters
		if utf8.RuneCountInString(countryISO2) != 2 {
			c.JSON(400, gin.H{"error": "countryISO2 must have 2 charakters"})
			return
		}
		var wanted_swift []models.SwiftCode

		fmt.Println("Searching for countryISO2:", countryISO2)

		result := database.DB.Where("country_ISO2 = ?", countryISO2).Find(&wanted_swift)
		if result.RowsAffected == 0 {
			c.JSON(404, gin.H{"error": "No Swift codes found for this country"})
			return
		}
		if result.Error != nil {
			c.JSON(500, gin.H{"error": "Database error"})
			return
		}

		response := responses.BuildEndpoint2Response(countryISO2, wanted_swift[0].CountryName, wanted_swift)
		c.JSON(200, response)
	})
}

func RegisterEndpoint3(r *gin.Engine) {
	r.POST("/v1/swift-codes", func(c *gin.Context) {
		var request responses.Endpoint3_request

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, responses.Message_response{
				Message: "Invalid request:" + err.Error(),
			})
			return
		}

		// check if countryiso2 has 2 charakters
		if utf8.RuneCountInString(request.CountryISO2) != 2 {
			c.JSON(400, responses.Message_response{
				Message: "countryISO2 must have 2 charakters",
			})
			return
		}

		// check if swiftcode has 11 charakters
		if utf8.RuneCountInString(request.SwiftCode) != 11 {
			c.JSON(400, responses.Message_response{
				Message: "swift code must have 11 characters",
			})
			return
		}

		// check if swift code already exist
		var existing models.SwiftCode
		result := database.DB.Where("swift_code = ?", request.SwiftCode).First(&existing)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				newSwift := models.SwiftCode{
					Address:       request.Address,
					BankName:      request.BankName,
					CountryISO2:   request.CountryISO2,
					CountryName:   request.CountryName,
					IsHeadquarter: request.IsHeadquarter,
					SwiftCode:     request.SwiftCode,
				}

				if err := database.DB.Create(&newSwift).Error; err != nil {
					c.JSON(500, responses.Message_response{
						Message: "Failed to create Swift code",
					})
					return
				}

				c.JSON(201, responses.Message_response{
					Message: "Swift code added successfully",
				})
				return
			}

			c.JSON(500, responses.Message_response{
				Message: "Database error",
			})
			return
		}
		c.JSON(400, responses.Message_response{
			Message: "swift code already exist",
		})
	})
}

func RegisterEndpoint4(r *gin.Engine) {
	r.DELETE("/v1/swift-codes/:swiftCode", func(c *gin.Context) {
		swiftCode := c.Param("swiftCode")
		if utf8.RuneCountInString(swiftCode) != 11 {
			c.JSON(400, responses.Message_response{
				Message: "swift code must have 11 characters",
			})
			return
		}

		var wanted_swift models.SwiftCode
		result := database.DB.Where("swift_code = ?", swiftCode).First(&wanted_swift)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(404, responses.Message_response{
					Message: "Swift code was not found",
				})
				return
			}
			c.JSON(500, responses.Message_response{
				Message: "Database error",
			})
			return
		}

		if err := database.DB.Unscoped().Delete(&wanted_swift).Error; err != nil {
			c.JSON(500, responses.Message_response{
				Message: "Failed to delete swift code",
			})
		}

		c.JSON(200, responses.Message_response{
			Message: "Swift code deleted successfully",
		})
	})
}
