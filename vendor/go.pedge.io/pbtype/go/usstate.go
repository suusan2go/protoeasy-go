package pbtype

import (
	"fmt"
	"strings"
)

// USStateCodeSimpleValueOf returns the value of the simple string s.
func USStateCodeSimpleValueOf(s string) (USStateCode, error) {
	value, ok := SimpleStringToUSStateCode[strings.ToUpper(s)]
	if !ok {
		return USStateCode_US_STATE_CODE_NONE, fmt.Errorf("pbtype: no USStateCode for %s", s)
	}
	return value, nil
}

// SimpleString returns the simple string.
func (c USStateCode) SimpleString() string {
	simpleString, ok := USStateCodeToSimpleString[c]
	if !ok {
		return ""
	}
	return simpleString
}
