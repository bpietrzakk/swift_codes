package responses

import (
	"github.com/bpietrzakk/swift_codes/internal/models"
)

// struct for headquarter (PROBLEM: branches show up as a center field)
type Endpoint1_Headquarter struct {
	Address string `json:"address"`
	BankName string `json:"bankName"`
	CountryISO2 string `json:"countryISO2"`
	CountryName string `json:"countryName"`
	IsHeadquarter bool `json:"isHeadquarter"`
	SwiftCode string `json:"swiftCode"`
	Branches []attached_swift `json:"branches"`
}

// struct for branches, which are attached to headquarter
type attached_swift struct {
	Address string `json:"address"`
	BankName string `json:"bankName"`
	CountryISO2 string `json:"countryISO2"`
	IsHeadquarter bool `json:"isHeadquarter"`
	SwiftCode string `json:"swiftCode"`
}

type Endpoint2 struct {
	CountryISO2 string `json:"countryISO2"`
	CountryName string `json:"countryName"`
	SwiftCodes []attached_swift `json:"swiftCodes"`
}

// // single branch response
// type single_branch struct {
// 	Address string `json:"address"`
// 	BankName string `json:"bankName"`
// 	CountryISO2 string `json:"countryISO2"`
// 	CountryName string `json:"countryName"`
// 	IsHeadquarter bool `json:"isHeadquarter"`
// 	SwiftCode string `json:"swiftCode"`
// }

func Build_HQB_response(headquarter_SC models.SwiftCode, branches []models.SwiftCode) Endpoint1_Headquarter {
	var branchList []attached_swift
	for _, b := range branches {
		branchList = append(branchList, attached_swift{
			Address: b.Address,
			BankName: b.BankName,
			CountryISO2: b.CountryISO2,
			IsHeadquarter: b.IsHeadquarter,
			SwiftCode: b.SwiftCode,
		})
	}
	return Endpoint1_Headquarter{
		Address:       headquarter_SC.Address,
		BankName:      headquarter_SC.BankName,
		CountryISO2:   headquarter_SC.CountryISO2,
		CountryName:   headquarter_SC.CountryName,
		IsHeadquarter: headquarter_SC.IsHeadquarter,
		SwiftCode:     headquarter_SC.SwiftCode,
		Branches:      branchList,
	}
}

func BuildEndpoint2Response(countryISO2, countryName string, swiftCodes []models.SwiftCode) Endpoint2 {
	var swiftCodesResponse []attached_swift
	for _, c := range swiftCodes {
		swiftCodesResponse = append(swiftCodesResponse, attached_swift{
			Address: c.Address,
			BankName: c.BankName,
			CountryISO2: c.CountryISO2,
			IsHeadquarter: c.IsHeadquarter,
			SwiftCode: c.SwiftCode,
		})
	}
	return Endpoint2{
		CountryISO2: countryISO2,
		CountryName: countryName,
		SwiftCodes: swiftCodesResponse,
	}
}

