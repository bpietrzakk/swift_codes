package responses

import (
	"github.com/bpietrzakk/swift_codes/internal/models"
)

// struct for headquarter (PROBLEM: branches show up as a center field)
type HeadquarterWithBranches struct {
	Address string `json:"address"`
	BankName string `json:"bankName"`
	CountryISO2 string `json:"countryISO2"`
	CountryName string `json:"countryName"`
	IsHeadquarter bool `json:"isHeadquarter"`
	SwiftCode string `json:"swiftCode"`
	Branches []attached_branch `json:"branches"`
}

// struct for branches, which are attached to headquarter
type attached_branch struct {
	Address string `json:"address"`
	BankName string `json:"bankName"`
	CountryISO2 string `json:"countryISO2"`
	IsHeadquarter bool `json:"isHeadquarter"`
	SwiftCode string `json:"swiftCode"`
}

// single branch response
type single_branch struct {
	Address string `json:"address"`
	BankName string `json:"bankName"`
	CountryISO2 string `json:"countryISO2"`
	CountryName string `json:"countryName"`
	IsHeadquarter bool `json:"isHeadquarter"`
	SwiftCode string `json:"swiftCode"`
}

func Build_HQB_response(headquarter_SC models.SwiftCode, branches []models.SwiftCode) HeadquarterWithBranches {
	var branchList []attached_branch
	for _, b := range branches {
		branchList = append(branchList, attached_branch{
			Address: b.Address,
			BankName: b.BankName,
			CountryISO2: b.CountryISO2,
			IsHeadquarter: b.IsHeadquarter,
			SwiftCode: b.SwiftCode,
		})
	}
	return HeadquarterWithBranches{
		Address:       headquarter_SC.Address,
		BankName:      headquarter_SC.BankName,
		CountryISO2:   headquarter_SC.CountryISO2,
		CountryName:   headquarter_SC.CountryName,
		IsHeadquarter: headquarter_SC.IsHeadquarter,
		SwiftCode:     headquarter_SC.SwiftCode,
		Branches:      branchList,
	}
}

