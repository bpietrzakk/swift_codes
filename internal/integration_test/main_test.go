package integrationtest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bpietrzakk/swift_codes/internal/database"
	"github.com/bpietrzakk/swift_codes/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
    "gorm.io/gorm"
	"github.com/bpietrzakk/swift_codes/internal/api"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	api.RegisterEndpoint1(r)
	api.RegisterEndpoint2(r)

	return r
}

func TestEndpoint1(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/AAISALTRXXX", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestEndpoint1_400error(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/TOOSHORT", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestEndpoint1_404error(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/NOTFOUNDXXX", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}

func TestEndpoint1_500error(t *testing.T) {
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/NOTFOUNDXXX", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}

func TestEndpoint2(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/country/AL", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}

func TestEndpoint2_400error(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/country/L", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 400, resp.Code)
}

func TestEndpoint2_404error(t *testing.T) {
	// set up database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
        t.Fatal(err)
    }

	
	database.DB = db
	db.AutoMigrate(&models.SwiftCode{})
	db.Create(&models.SwiftCode{
        CountryISO2: "AL",
		SwiftCode: "AAISALTRXXX",
		CodeType: "BIC11",
		BankName: "UNITED BANK OF ALBANIA SH.A",
		Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
		TownName: "TIRANA",
		CountryName: "ALBANIA",
		TimeZone: "Europe/Tirane",
		IsHeadquarter: true,
    })
	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/country/NO", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}

// Test below works in vs code gui, but dont work in terminal 

func TestEndpoint2_500error(t *testing.T) {

	router := SetupRouter()
	req, _ := http.NewRequest("GET", "/v1/swift-codes/country/SL", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
}