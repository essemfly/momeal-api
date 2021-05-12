package utils

import (
	"fmt"
	"strconv"
	"unicode"
)

// CheckErr function
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ParsePriceString(letter string) int {
	ret := ""
	for _, val := range letter {
		if unicode.IsDigit(val) {
			ret += string(val)
		}
	}
	retInt, _ := strconv.Atoi(ret)
	return retInt
}
