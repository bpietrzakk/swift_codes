package api

import (
	"errors"
	"unicode/utf8"

	"github.com/bpietrzakk/swift_codes/internal/database"
	"github.com/bpietrzakk/swift_codes/internal/models"
	"github.com/bpietrzakk/swift_codes/internal/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ENDPOINT 1: GET: /v1/swift-codes/{swift-code}
//	response for hq with branches
//	response for branch single
func RegisterEndpoint1(r *gin.Engine) {
	r.GET("/v1/swift-codes/:swift_code", func(c *gin.Context) {
		swift_code := c.Param("swift_code")
		// sprawdzic czy swiftcode ma 11 znakow
		if utf8.RuneCountInString(swift_code) != 11 {
			c.JSON(400, gin.H{"error": "swift code must have 11 characters"})
            return
		}
		// przeszukac swiftcode w bazie danych
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
				"address": wanted_swift.Address,
				"bankName": wanted_swift.BankName,
				"countryISO2": wanted_swift.CountryISO2,
				"countryName": wanted_swift.CountryName,
				"isHeadquarter": wanted_swift.IsHeadquarter,
				"swiftCode": wanted_swift.SwiftCode,
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

// func RegisterEndpoint2(r *gin.Engine){
// 	r.GET("/v1/swift-codes/country/:countryISO2code", func(c *gin.Context) {

// 	})
// }