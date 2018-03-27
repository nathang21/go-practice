package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	tables := []struct {
		have string
		want string
		err  error
	}{
		{"MXXVIII", "ⅯⅩⅩⅤⅠⅠⅠ", nil},
		{"XIlLLbCM", "", errors.New("unable to match string to roman numeral unicode")},
		{"", "", errors.New("incorrect use of subractive notation")},
	}
	for _, test := range tables {
		fmt.Println(convert(test.have))
		if out, err := convert(test.have); out != test.want {
			t.Errorf("error during unicode conversion: %s.", err)
		}
	}
}
