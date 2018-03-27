// Converts string input to valid Roman Numeral unicode
package main

import "fmt"
import "strings"

var uni = map[string]string{
	"I": "Ⅰ",
	"V": "Ⅴ",
	"X": "Ⅹ",
	"L": "Ⅼ",
	"C": "Ⅽ",
	"D": "Ⅾ",
	"M": "Ⅿ",
}

func convert(s string) (out string, err error) {
	r := strings.ToUpper(s)
	for _, v := range r {
		val, ok := uni[string(v)]
		if ok {
			out += string(val)
			continue
		}
		return "", fmt.Errorf("unable to match rune %s to roman numeral unicode", val)
	}
	return out, nil
}

func main() {
	fmt.Println("MXXVIII")
	fmt.Println(convert("MXXVIII"))
	fmt.Println(convert("XIlLLbCM"))
}
