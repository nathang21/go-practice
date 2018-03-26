package main

import (
	"errors"
	"testing"
)

func TestGetStaticRomanNums(t *testing.T) {
	tables := []struct {
		have string
		want int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},

		{"D", 500},
		{"M", 1000},
	}
	for _, test := range tables {
		if out, _ := GetStaticRomanNums(test.have); out != test.want {
			t.Errorf("Base value is incorrect")
		}
	}
}

func TestSubractiveNotation(t *testing.T) {
	tables := []struct {
		have string
		want string
		err  error
	}{
		{"IV", "IIII", nil},
		{"IX", "VIIII", nil},
		{"XL", "XXXX", nil},
		{"XC", "LXXXX", nil},
		{"CD", "CCCC", nil},
		{"CM", "DCCCC", nil},
		{"MMMCMXCIX", "", errors.New("incorrect use of subractive notation")},
		{"", "", errors.New("Incorrect use of subractive notation")},
	}
	for _, test := range tables {
		if out, err := SubractiveNotation(test.have); out != test.want {
			t.Errorf("Error during notation conversion: %s.", err)
		}
	}
}

func TestReplace(t *testing.T) {
	tables := []struct {
		have string
		want string
	}{
		{"IV", "IIII"},
		{"IX", "VIIII"},
		{"XL", "XXXX"},
		{"XC", "LXXXX"},
		{"CD", "CCCC"},
		{"CM", "DCCCC"},
		{"MMMCMXCIX", "MMMDCCCCLXXXXVIIII"},
		{"V", "V"},
		{"MXXIII", "MXXIII"},
	}
	for _, test := range tables {
		if out := Replace(test.have); out != test.want {
			t.Errorf("The replacement applied was incorrect.")
		}
	}
}

func TestValidate(t *testing.T) {
	tables := []struct {
		have        string
		want        error
		ExpectError bool
	}{
		{"IXV", errors.New("String is not valid roman numeral notation"), true},
		{"MDI", nil, false},
		{"LIV", nil, false},
		{"IV", nil, false},
		{"IX", nil, false},
		{"XCVII", nil, false},
		{"DCCCXCI", nil, false},
		{"DCCXCVIII", nil, false},
		{"MIV", nil, false},
		{"M", nil, false},
		{"MVI", nil, false},
		{"MXXIII", nil, false},
		{"MMXIV", nil, false},
		{"MMMCMXCIX", nil, false},
		{"IVL", errors.New("String is not valid roman numeral notation"), true},
		{"MVM", errors.New("String is not valid roman numeral notation"), true},
		{"CDM", errors.New("String is not valid roman numeral notation"), true},
	}
	for _, test := range tables {
		err := Validate(test.have)
		if test.ExpectError && err == nil {
			t.Errorf("Did not properly determine in-validity of roman numeral notation, expected Error.")
		}
		if !test.ExpectError && err != nil {
			t.Errorf("Did not expect error for %s, but got: %s.", test.have, err)
		}
	}
}

func TestCalcNotation(t *testing.T) {
	tables := []struct {
		have string
		want int
	}{
		{"IXV", 0},
		{"MDI", 1501},
		{"LIV", 54},
		{"IV", 4},
		{"IX", 9},
		{"IXV", 0},
		{"XCVII", 97},
		{"DCCCXCI", 891},
		{"DCCXCVIII", 798},
		{"MIV", 1004},
		{"M", 1000},
		{"MVI", 1006},
		{"MXXIII", 1023},
		{"MMXIV", 2014},
		{"MMMCMXCIX", 3999},
		{"IVL", 0},
		{"MVM", 0},
		{"CDM", 0},
	}
	for _, test := range tables {
		if out, _ := CalcNotation(test.have); out != test.want {
			t.Errorf("Roman numeral conversion to decimal is incorrect. %d does not equal %d.", out, test.want)
		}
	}
}
