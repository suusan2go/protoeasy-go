package pbtype // import "go.pedge.io/pbtype/go"

import (
	"fmt"
	"math"
	"strings"
)

// CountryAlpha2CodeSimpleValueOf returns the value of the simple string s.
func CountryAlpha2CodeSimpleValueOf(s string) (CountryAlpha2Code, error) {
	value, ok := SimpleStringToCountryAlpha2Code[strings.ToUpper(s)]
	if !ok {
		return CountryAlpha2Code_COUNTRY_ALPHA_2_CODE_NONE, fmt.Errorf("pbtype: no CountryAlpha2Code for %s", s)
	}
	return value, nil
}

// CountryAlpha3CodeSimpleValueOf returns the value of the simple string s.
func CountryAlpha3CodeSimpleValueOf(s string) (CountryAlpha3Code, error) {
	value, ok := SimpleStringToCountryAlpha3Code[strings.ToUpper(s)]
	if !ok {
		return CountryAlpha3Code_COUNTRY_ALPHA_3_CODE_NONE, fmt.Errorf("pbtype: no CountryAlpha3Code for %s", s)
	}
	return value, nil
}

// CurrencyCodeSimpleValueOf returns the value of the simple string s.
func CurrencyCodeSimpleValueOf(s string) (CurrencyCode, error) {
	value, ok := SimpleStringToCurrencyCode[strings.ToUpper(s)]
	if !ok {
		return CurrencyCode_CURRENCY_CODE_NONE, fmt.Errorf("pbtype: no CurrencyCode for %s", s)
	}
	return value, nil
}

// SimpleString returns the simple string.
func (c CountryAlpha2Code) SimpleString() string {
	simpleString, ok := CountryAlpha2CodeToSimpleString[c]
	if !ok {
		return ""
	}
	return simpleString
}

// Country returns the associated Country, or nil if no Country is known.
func (c CountryAlpha2Code) Country() *Country {
	country, ok := CountryAlpha2CodeToCountry[c]
	if !ok {
		return nil
	}
	return country
}

// Currency returns the associated Currency, or nil if no Country or Currency is known.
func (c CountryAlpha2Code) Currency() *Currency {
	country := c.Country()
	if country == nil {
		return nil
	}
	currency, ok := CurrencyCodeToCurrency[country.CurrencyCode]
	if !ok {
		return nil
	}
	return currency
}

// SimpleString returns the simple string.
func (c CountryAlpha3Code) SimpleString() string {
	simpleString, ok := CountryAlpha3CodeToSimpleString[c]
	if !ok {
		return ""
	}
	return simpleString
}

// Country returns the associated Country, or nil if no Country is known.
func (c CountryAlpha3Code) Country() *Country {
	country, ok := CountryAlpha3CodeToCountry[c]
	if !ok {
		return nil
	}
	return country
}

// Currency returns the associated Currency, or nil if no Country or Currency is known.
func (c CountryAlpha3Code) Currency() *Currency {
	country := c.Country()
	if country == nil {
		return nil
	}
	currency, ok := CurrencyCodeToCurrency[country.CurrencyCode]
	if !ok {
		return nil
	}
	return currency
}

// SimpleString returns the simple string.
func (c CurrencyCode) SimpleString() string {
	simpleString, ok := CurrencyCodeToSimpleString[c]
	if !ok {
		return ""
	}
	return simpleString
}

// Currency returns the associated Currency, or nil if no Currency is known.
func (c CurrencyCode) Currency() *Currency {
	currency, ok := CurrencyCodeToCurrency[c]
	if !ok {
		return nil
	}
	return currency
}

func (c CurrencyCode) minorMultiplier() int64 {
	// TODO(pedge): will panic if c.Currency() == nil
	return int64(math.Pow(10, float64(6-c.Currency().MinorUnit)))
}
