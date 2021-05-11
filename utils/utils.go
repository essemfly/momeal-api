package utils

import (
	"fmt"
)

// CheckErr function
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
