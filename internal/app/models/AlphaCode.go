package models

import (
	"fmt"
)

type AlphaCodeMass []AlphaCode

type AlphaCode struct {
	AlphaTwo        string `json:"alpha_two"`
	NumericalCode   int    `json:"numerical_code"`
	FullCountryName string `json:"full_country_name"`
}

func SearchCode(inputCode string) error {
	for _, v := range *alphaCodeGet() {
		if inputCode == v.AlphaTwo {
			return nil
		}
	}
	return fmt.Errorf("данный код страны отсутствует в системе")
}

func FullCountryNameSMS(input []SMSData) {
	for i, data := range input {
		input[i].Country = ChangeCountryName(data.Country)
	}
}

func FullCountryNameMMS(input []MMSData) {
	for i, data := range input {
		input[i].Country = ChangeCountryName(data.Country)
	}
}

func ChangeCountryName(inputCode string) (result string) {
	for _, v := range *alphaCodeGet() {
		if inputCode == v.AlphaTwo {
			result = v.FullCountryName
		}
	}
	return
}

func alphaCodeGet() *AlphaCodeMass {
	//ISO 3166-2 and ./skillbox-diploma/emulator.go.288
	return &AlphaCodeMass{
		{
			AlphaTwo:        "RU",
			NumericalCode:   643,
			FullCountryName: "Russia",
		},
		{
			AlphaTwo:        "US",
			NumericalCode:   840,
			FullCountryName: "USA",
		},
		{
			AlphaTwo:        "GB",
			NumericalCode:   826,
			FullCountryName: "Great Britain",
		},
		{
			AlphaTwo:        "FR",
			NumericalCode:   250,
			FullCountryName: "France",
		},
		{
			AlphaTwo:        "BL",
			NumericalCode:   652,
			FullCountryName: "Saint Barthelemy",
		},
		{
			AlphaTwo:        "AT",
			NumericalCode:   040,
			FullCountryName: "Austria",
		},
		{
			AlphaTwo:        "BG",
			NumericalCode:   100,
			FullCountryName: "Bulgaria",
		},
		{
			AlphaTwo:        "DK",
			NumericalCode:   208,
			FullCountryName: "Denmark",
		},
		{
			AlphaTwo:        "CA",
			NumericalCode:   124,
			FullCountryName: "Canada",
		},
		{
			AlphaTwo:        "ES",
			NumericalCode:   724,
			FullCountryName: "Spain",
		},
		{
			AlphaTwo:        "CH",
			NumericalCode:   756,
			FullCountryName: "Switzerland",
		},
		{
			AlphaTwo:        "TR",
			NumericalCode:   792,
			FullCountryName: "Turkey",
		},
		{
			AlphaTwo:        "PE",
			NumericalCode:   604,
			FullCountryName: "Peru",
		},
		{
			AlphaTwo:        "NZ",
			NumericalCode:   554,
			FullCountryName: "New Zealand",
		},
		{
			AlphaTwo:        "MC",
			NumericalCode:   492,
			FullCountryName: "Monaco",
		},
	}
}
