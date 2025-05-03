package models

type SwiftCode struct {
	ID 				uint	`gorm:"primaryKey" json:"id"`
	CountryISO2		string	`gorm:"column:country_iso2;type:varchar(2);not null" json:"CountryISO2"`
	SwiftCode		string	`gorm:"column:swift_code;type:varchar(11);unique;not null;index" json:"SwiftCode"`
	CodeType		string	`gorm:"column:code_type;type:varchar(5)" json:"CodeType"`
	BankName		string	`gorm:"column:bank_name;type:text" json:"BankName"`
	Address			string	`gorm:"column:address;type:text" json:"Address"`
	TownName		string	`gorm:"column:town_name;type:varchar(30)" json:"TownName"`
	CountryName		string	`gorm:"column:country_name;type:varchar(30)" json:"CountryName"`
	TimeZone		string	`gorm:"column:time_zone;type:varchar(50)" json:"TimeZone"`
	IsHeadquarter	bool	`gorm:"column:is_head_quarter;default:false" json:"IsHeadquarter"`
}