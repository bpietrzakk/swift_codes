package parser

import (
	"fmt"
	"encoding/csv"
	"os"
	"strings"
	"github.com/bpietrzakk/swift_codes/internal/models"
)

func ParseSwiftCodesCSV(filepath string) ([]models.SwiftCode, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var  swiftCodes []models.SwiftCode
	header:=records[0]

	if len(header) < 8 {
		return nil, fmt.Errorf("invalid CSV header")
	}

	for _, record := range records[1:] { // Start from the second row

		swiftCode := models.SwiftCode{
			CountryISO2:	record[0],
			SwiftCode: record[1],
			CodeType: record[2],
			BankName: record[3],
			Address: record[4],
			TownName: record[5],
			CountryName: record[6],
			TimeZone: record[7],
			IsHeadquarter: strings.HasSuffix(record[1], "XXX"),
		}
		swiftCodes = append(swiftCodes, swiftCode)
	}
	return swiftCodes, nil
}