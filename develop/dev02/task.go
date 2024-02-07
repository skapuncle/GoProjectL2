package main

import (
	"fmt"
	"strconv"
	"strings"
)

func unpack(s string) (string, error) {
	var res = strings.Builder{}
	digitBuffer := strings.Builder{}

	for _, char := range s {
		switch {
		case char >= '0' && char <= '9':
			digitBuffer.WriteRune(char)
		default:
			repeatTimes, _ := strconv.Atoi(digitBuffer.String())
			if repeatTimes == 0 {
				repeatTimes = 1
			}
			for i := 0; i < repeatTimes; i++ {
				res.WriteRune(char)
			}
			digitBuffer.Reset()
		}
	}

	if digitBuffer.Len() > 0 {
		return "", fmt.Errorf("invalid string")
	}

	return res.String(), nil
}
func main() {
	fmt.Println(unpack(`a4bc2d5e`))
	fmt.Println(unpack(`abcd`))
	fmt.Println(unpack(`45`))
	fmt.Println(unpack(``))
}
