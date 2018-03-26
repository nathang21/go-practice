// Validates and converts roman numeral notation to decimal
package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
var conversions = []string{"IV", "IX", "XL", "XC", "CD", "CM"}

// GetStaticRomanNums takes a Roman numeral symbol and returns the equivalent decimal value
func GetStaticRomanNums(sym string) (int, error) {
	if val, ok := m[sym]; ok {
		return val, nil
	}
	return -1, errors.New("not a valid roman numeral character")
}

// SubractiveNotation takes in a subtractive notation form and returns the full expanded form
func SubractiveNotation(sym string) (string, error) {
	switch sym {
	case "IV":
		return "IIII", nil
	case "IX":
		return "VIIII", nil
	case "XL":
		return "XXXX", nil
	case "XC":
		return "LXXXX", nil
	case "CD":
		return "CCCC", nil
	case "CM":
		return "DCCCC", nil
	}
	return "", errors.New("incorrect use of subractive notation")
}

// Replace takes in a string to check if substrings need to be converted and outupts new string
func Replace(sym string) string {
	for _, c := range conversions {
		if convert, err := SubractiveNotation(c); err == nil {
			sym = strings.Replace(sym, c, convert, -1)
		}
	}
	return sym
}

// Validate takes in a string and returns a boolean if format is correct.
func Validate(sym string) error {
	re := regexp.MustCompile("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")
	if re.MatchString(sym) != true {
		return errors.New("String is not valid roman numeral notation")
	}
	return nil
}

// CalcNotation takes a roman numeral string and outputs the sum
func CalcNotation(sym string) (int, error) {
	if err := Validate(sym); err != nil {
		return 0, err
	}
	long := Replace(sym)
	total := 0
	for _, c := range long {
		t, _ := GetStaticRomanNums(string(c))
		total += t
	}
	return total, nil
}

func main() {
	fmt.Println(Validate("IVL"))
	fmt.Println(Validate("MDI"))
	fmt.Println(Validate("LIV"))
	fmt.Println(CalcNotation("LIV"))
	fmt.Println(Replace("LIV"))
	fmt.Println(Validate("IV"))
	fmt.Println(Validate("XCVII"))
	fmt.Println(CalcNotation("DCCCXCI"))
	fmt.Println(CalcNotation("DCCXCVIII"))
	fmt.Println(Validate("MIV"))
	fmt.Println(CalcNotation("M"))
	fmt.Println(CalcNotation("MVI"))
	fmt.Println(CalcNotation("MXXIII"))
	fmt.Println(CalcNotation("MMXIV"))
	fmt.Println(Replace("MXXIII"))
}
