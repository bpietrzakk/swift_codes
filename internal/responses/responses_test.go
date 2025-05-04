package responses

import (
	"reflect"
	"testing"
	"github.com/bpietrzakk/swift_codes/internal/models"
)

func TestBuild_HQB_response(t *testing.T) {
	headquarter := models.SwiftCode{
		Address:       "HQ Address",
		BankName:      "Test Bank",
		CountryISO2:   "PL",
		CountryName:   "Poland",
		IsHeadquarter: true,
		SwiftCode:     "HQSWIFT12345",
	}

	branches := []models.SwiftCode{
		{
			Address:       "Branch 1",
			BankName:      "Test Bank",
			CountryISO2:   "PL",
			IsHeadquarter: false,
			SwiftCode:     "BRANCH12345",
		},
	}

	expected := Endpoint1_Headquarter{
		Address:       "HQ Address",
		BankName:      "Test Bank",
		CountryISO2:   "PL",
		CountryName:   "Poland",
		IsHeadquarter: true,
		SwiftCode:     "HQSWIFT12345",
		Branches: []AttachedSwift{
			{
				Address:       "Branch 1",
				BankName:      "Test Bank",
				CountryISO2:   "PL",
				IsHeadquarter: false,
				SwiftCode:     "BRANCH12345",
			},
		},
	}

	result := Build_HQB_response(headquarter, branches)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestBuild_HQB_response_EmptyBranches(t *testing.T) {
	headquarter := models.SwiftCode{
		Address:       "HQ Address",
		BankName:      "Test Bank",
		CountryISO2:   "PL",
		CountryName:   "Poland",
		IsHeadquarter: true,
		SwiftCode:     "HQSWIFT12345",
	}

	branches := []models.SwiftCode{} // empty

	expected := Endpoint1_Headquarter{
		Address:       "HQ Address",
		BankName:      "Test Bank",
		CountryISO2:   "PL",
		CountryName:   "Poland",
		IsHeadquarter: true,
		SwiftCode:     "HQSWIFT12345",
		Branches:      nil, // empty slice
	}

	result := Build_HQB_response(headquarter, branches)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestBuildEndpoint2Response(t *testing.T) {

	CountryISO2 := "PL"
	CountryName := "POLAND"
	swiftCodes := []models.SwiftCode{
		{
			CountryISO2:   "PL",
			SwiftCode:     "HQSWIFT12XXX",
			BankName:      "Test Bank 1",
			Address:       "Test 1",
			IsHeadquarter: true,
		},
		{
			CountryISO2:   "PL",
			SwiftCode:     "BRSWIFT12345",
			BankName:      "Test Bank 2",
			Address:       "Test 2",
			IsHeadquarter: false,
		},
	}
	expected := Endpoint2{
		CountryISO2: "PL",
		CountryName: "POLAND",
		SwiftCodes: []AttachedSwift{
			{
				CountryISO2:   "PL",
				SwiftCode:     "HQSWIFT12XXX",
				BankName:      "Test Bank 1",
				Address:       "Test 1",
				IsHeadquarter: true,
			},
			{
				CountryISO2:   "PL",
				SwiftCode:     "BRSWIFT12345",
				BankName:      "Test Bank 2",
				Address:       "Test 2",
				IsHeadquarter: false,
			},
		},
	}
	result := BuildEndpoint2Response(CountryISO2, CountryName, swiftCodes)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}
