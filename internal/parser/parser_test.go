package parser

import (
	"testing"
	"reflect"
	"strings"
	"github.com/bpietrzakk/swift_codes/internal/models"
)

func TestParseSwiftCodesCSV(t *testing.T) {
	csvPath := "../data/test/test.csv"

	result, err := ParseSwiftCodesCSV(csvPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []models.SwiftCode{
		{
			CountryISO2: "AL",
			SwiftCode: "AAISALTRXXX",
			CodeType: "BIC11",
			BankName: "UNITED BANK OF ALBANIA SH.A",
			Address: "HYRJA 3 RR. DRITAN HOXHA ND. 11 TIRANA, TIRANA, 1023",
			TownName: "TIRANA",
			CountryName: "ALBANIA",
			TimeZone: "Europe/Tirane",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "BG",
			SwiftCode: "ABIEBGS1XXX",
			CodeType: "BIC11",
			BankName: "ABV INVESTMENTS LTD",
			Address: "TSAR ASEN 20  VARNA, VARNA, 9002",
			TownName: "VARNA",
			CountryName: "BULGARIA",
			TimeZone: "Europe/Sofia",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "BG",
			SwiftCode: "ADCRBGS1XXX",
			CodeType: "BIC11",
			BankName: "ADAMANT CAPITAL PARTNERS AD",
			Address: "JAMES BOURCHIER BLVD 76A HILL TOWER SOFIA, SOFIA, 1421",
			TownName: "SOFIA",
			CountryName: "BULGARIA",
			TimeZone: "Europe/Sofia",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "UY",
			SwiftCode: "AFAAUYM1XXX",
			CodeType: "BIC11",
			BankName: "AFINIDAD A.F.A.P.S.A.",
			Address: "PLAZA INDEPENDENCIA 743  MONTEVIDEO, MONTEVIDEO, 11000",
			TownName: "MONTEVIDEO",
			CountryName: "URUGUAY",
			TimeZone: "America/Montevideo",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "MC",
			SwiftCode: "AGRIMCM1XXX",
			CodeType: "BIC11",
			BankName: "CREDIT AGRICOLE MONACO (CRCA PROVENCE COTE D'AZUR MONACO)",
			Address: "23 BOULEVARD PRINCESSE CHARLOTTE  MONACO, MONACO, 98000",
			TownName: "MONACO",
			CountryName: "MONACO",
			TimeZone: "Europe/Monaco",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "PL",
			SwiftCode: "AIPOPLP1XXX",
			CodeType: "BIC11",
			BankName: "SANTANDER CONSUMER BANK SPOLKA AKCYJNA",
			Address: "STRZEGOMSKA 42C  WROCLAW, DOLNOSLASKIE, 53-611",
			TownName: "WROCLAW",
			CountryName: "POLAND",
			TimeZone: "Europe/Warsaw",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "LV",
			SwiftCode: "AIZKLV22XXX",
			CodeType: "BIC11",
			BankName: "ABLV BANK, AS IN LIQUIDATION",
			Address: "MIHAILA TALA STREET 1  RIGA, RIGA, LV-1045",
			TownName: "RIGA",
			CountryName: "LATVIA",
			TimeZone: "Europe/Riga",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "MT",
			SwiftCode: "AKBKMTMTXXX",
			CodeType: "BIC11",
			BankName: "AKBANK T.A.S. (MALTA BRANCH)",
			Address: "FLOOR 6, PORTOMASO BUSINESS TOWER 01 PORTOMASO PTM - ST. JULIAN'S ST. JULIAN'S, STJ 4011",
			TownName: "ST. JULIAN'S",
			CountryName: "MALTA",
			TimeZone: "Europe/Malta",
			IsHeadquarter: true,
		},
		{
			CountryISO2: "PL",
			SwiftCode: "ALBPPLP1BMW",
			CodeType: "BIC11",
			BankName: "ALIOR BANK SPOLKA AKCYJNA",
			Address: "  WARSZAWA, MAZOWIECKIE",
			TownName: "WARSZAWA",
			CountryName: "POLAND",
			TimeZone: "Europe/Warsaw",
			IsHeadquarter: false,
		},
		{
			CountryISO2: "PL",
			SwiftCode: "ALBPPLPWXXX",
			CodeType: "BIC11",
			BankName: "ALIOR BANK SPOLKA AKCYJNA",
			Address: "LOPUSZANSKA BUSINESS PARK LOPUSZANSKA 38 D WARSZAWA, MAZOWIECKIE, 02-232",
			TownName: "WARSZAWA",
			CountryName: "POLAND",
			TimeZone: "Europe/Warsaw",
			IsHeadquarter: true,
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %+v, got: %+v", expected, result)
	}
}

func TestParseSwiftCodesCSV_Empty(t *testing.T) {
	csvPath := "../data/test/empty_test.csv"

	expextedError := "failed to read CSV: "
	_, err := ParseSwiftCodesCSV(csvPath)
	if err == nil {
		t.Fatalf("Expected error: %q", expextedError)
	}
	if !strings.Contains(err.Error(), expextedError){
		t.Errorf("Expected error: %q, got: %q", expextedError, err.Error())
	}
}

func TestParseSwiftCodesCSV_InvalidHeader(t *testing.T){
	csvPath := "../data/test/invalidHeader_test.csv"

	expextedError := "invalid CSV header"
	_, err := ParseSwiftCodesCSV(csvPath)
	if err == nil {
		t.Fatalf("Expected error: %q", expextedError)
	}
	if !strings.Contains(err.Error(), expextedError){
		t.Errorf("Expected error: %q, got: %q", expextedError, err.Error())
	}
}

func TestParseSwiftCodesCSV_UnexistingFile(t *testing.T){
	csvPath := "../data/test/UnExistingFile.csv"

	expextedError := "failed to open file: "
	_, err := ParseSwiftCodesCSV(csvPath)
	if err == nil {
		t.Fatalf("Expected error: %q", expextedError)
	}
	if !strings.Contains(err.Error(), expextedError){
		t.Errorf("Expected error: %q, got: %q", expextedError, err.Error())
	}
}
