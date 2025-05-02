package models

type SwiftCode struct {
	ID 				uint	`gorm:"primaryKey" json:"id"`
	CountryISO2		string	`gorm:"type:varchar(2);not null" json:"CountryISO2"`
	SwiftCode		string	`gorm:"type:varchar(11);unique;not null;index" json:"SwiftCode"`
	CodeType		string	`gorm:"type:varchar(5)" json:"CodeType"`
	BankName		string	`gorm:"type:text" json:"BankName"`
	Address			string	`gorm:"type:text" json:"Address"`
	TownName		string	`gorm:"type:varchar(30)" json:"TownName"`
	CountryName		string	`gorm:"type:varchar(30)" json:"CountryName"`
	TimeZone		string	`gorm:"type:varchar(50)" json:"TimeZone"`
	IsHeadquarter	bool	`gorm:"default:false" json:"IsHeadquarter"`
}