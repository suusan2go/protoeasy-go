package pbtype

import (
	"fmt"

	"go.pedge.io/googleapis/google/type"
)

// MoneyMather performs math operations on Money.
type MoneyMather interface {
	Plus(MoneyMather) MoneyMather
	Minus(MoneyMather) MoneyMather
	Times(MoneyMather) MoneyMather
	Div(MoneyMather) MoneyMather
	Abs() MoneyMather
	Result() (*Money, error)

	noCurrencyCode() bool
	currencyCode() CurrencyCode
	// TODO(pedge): precise and imprecise modes
	valueMicros() float64
	errors() []error
}

// NewMoneyMather returns a new MoneyMather that will directly use the valueMicros and not check CurrencyCodes.
func NewMoneyMather(valueMicros float64) MoneyMather {
	return newMoneyMather(valueMicros)
}

// NewMoney returns a new Money for the given CurrencyCode and valueMicros.
func NewMoney(currencyCode CurrencyCode, valueMicros int64) *Money {
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  valueMicros,
	}
}

// NewMoneyUSD returns a new Money for USD for the given valueMicros.
func NewMoneyUSD(valueMicros int64) *Money {
	return NewMoney(CurrencyCodeUSD, valueMicros)
}

// NewMoneyEUR returns a new Money for EUR for the given valueMicros.
func NewMoneyEUR(valueMicros int64) *Money {
	return NewMoney(CurrencyCodeEUR, valueMicros)
}

// NewMoneySimple returns a new Money for the given CurrencyCode and value.
//
// ValueMinorUnits will use the units of the CurrencyCode.
// If valueUnits is negative, valueMinorUnits will be converted to negative.
func NewMoneySimple(currencyCode CurrencyCode, valueUnits int64, valueMinorUnits int64) *Money {
	if valueUnits < 0 {
		valueMinorUnits = -valueMinorUnits
	}
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  unitsAndMicroPartToMicros(valueUnits, valueMinorUnits*currencyCode.minorMultiplier()),
	}
}

// NewMoneySimpleUSD returns a new Money for USD for the given and value.
//
// If valueDollars is negative, valueCents will be converted to negative.
func NewMoneySimpleUSD(valueDollars int64, valueCents int64) *Money {
	return NewMoneySimple(CurrencyCodeUSD, valueDollars, valueCents)
}

// NewMoneySimpleEUR returns a new Money for EUR for the given and value.
//
// If valueEuros is negative, valueCents will be converted to negative.
func NewMoneySimpleEUR(valueEuros int64, valueCents int64) *Money {
	return NewMoneySimple(CurrencyCodeEUR, valueEuros, valueCents)
}

// GoogleMoneyToMoney converts a google.type.Money to Money.
func GoogleMoneyToMoney(googleMoney *google_type.Money) (*Money, error) {
	currencyCode, err := CurrencyCodeSimpleValueOf(googleMoney.CurrencyCode)
	if err != nil {
		return nil, err
	}
	return &Money{
		CurrencyCode: currencyCode,
		ValueMicros:  unitsAndNanoPartToMicros(googleMoney.Units, googleMoney.Nanos),
	}, nil
}

// ToGoogleMoney converts the given Money into a google.type.Money
func (m *Money) ToGoogleMoney() *google_type.Money {
	units, nanoPart := microsToUnitsAndNanoPart(m.ValueMicros)
	return &google_type.Money{
		CurrencyCode: m.CurrencyCode.SimpleString(),
		Units:        units,
		Nanos:        nanoPart,
	}
}

// Math returns a new MoneyMather for the Money.
func (m *Money) Math() MoneyMather {
	return newMoneyMatherForMoney(m)
}

// SimpleString returns the simple string for the Money.
func (m *Money) SimpleString() string {
	return fmt.Sprintf("%d.%06d", m.ValueMicros/1000000, m.ValueMicros%1000000)
}

type moneyMather struct {
	nocc bool
	cc   CurrencyCode
	vm   float64
	errs []error
}

func newMoneyMatherForMoney(money *Money) *moneyMather {
	return &moneyMather{
		cc: money.CurrencyCode,
		vm: float64(money.ValueMicros),
	}
}

func newMoneyMather(valueMicros float64) *moneyMather {
	return &moneyMather{
		nocc: true,
		vm:   valueMicros,
	}
}

func (m *moneyMather) Plus(moneyMather MoneyMather) MoneyMather {
	if !m.ok(moneyMather) {
		return m
	}
	m.vm += moneyMather.valueMicros()
	return m
}

func (m *moneyMather) Minus(moneyMather MoneyMather) MoneyMather {
	if !m.ok(moneyMather) {
		return m
	}
	m.vm -= moneyMather.valueMicros()
	return m
}

func (m *moneyMather) Times(moneyMather MoneyMather) MoneyMather {
	if !m.ok(moneyMather) {
		return m
	}
	m.vm *= moneyMather.valueMicros()
	return m
}

func (m *moneyMather) Div(moneyMather MoneyMather) MoneyMather {
	if !m.ok(moneyMather) {
		return m
	}
	if moneyMather.valueMicros() == 0 {
		m.addErrors(fmt.Errorf("pbtype: cannot divide by 0"))
		return m
	}
	m.vm /= moneyMather.valueMicros()
	return m
}

func (m *moneyMather) Abs() MoneyMather {
	if !m.ok(nil) {
		return m
	}
	if m.vm < 0 {
		m.vm = -m.vm
	}
	return m
}

func (m *moneyMather) Result() (*Money, error) {
	if len(m.errs) > 0 {
		return nil, fmt.Errorf("%v", m.errs)
	}
	return &Money{
		CurrencyCode: m.cc,
		ValueMicros:  int64(m.vm),
	}, nil
}

func (m *moneyMather) noCurrencyCode() bool {
	return m.nocc
}

func (m *moneyMather) currencyCode() CurrencyCode {
	return m.cc
}

func (m *moneyMather) valueMicros() float64 {
	return m.vm
}

func (m *moneyMather) errors() []error {
	return m.errs
}

func (m *moneyMather) ok(moneyMather MoneyMather) bool {
	errs := moneyMather.errors()
	if len(errs) > 0 {
		m.addErrors(errs...)
		return false
	}
	if moneyMather != nil {
		if !m.nocc && !moneyMather.noCurrencyCode() && m.cc != moneyMather.currencyCode() {
			m.addErrors(fmt.Errorf("pbtype: mismatched CurrencyCodes: %s %s", m.cc, moneyMather.currencyCode()))
			return false
		}
		if m.nocc && !moneyMather.noCurrencyCode() {
			m.nocc = false
			m.cc = moneyMather.currencyCode()
		}
	}
	return true
}

func (m *moneyMather) addErrors(errs ...error) {
	m.errs = append(m.errs, errs...)
}

func unitsAndMicroPartToMicros(units int64, micros int64) int64 {
	return unitsToMicros(units) + micros
}

func unitsAndNanoPartToMicros(units int64, nanos int32) int64 {
	return unitsToMicros(units) + int64(nanos/1000)
}

func microsToUnitsAndMicroPart(micros int64) (int64, int64) {
	return micros / 1000000, micros % 1000000
}

func microsToUnitsAndNanoPart(micros int64) (int64, int32) {
	return micros / 1000000, int32(micros%1000000) * 1000
}

func unitsToMicros(units int64) int64 {
	return units * 1000000
}
