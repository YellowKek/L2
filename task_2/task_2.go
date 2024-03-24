package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	var s = "a4bc2d5e"
	fmt.Scan(&s)
	res, err := Unpack(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}

func Unpack(s string) (string, error) {
	if validate(s) {
		return printString(s), nil
	} else {
		return "", errors.New("некорректная строка")
	}
}

func validate(s string) bool {
	var digitFlag = false
	for _, c := range s {
		if unicode.IsDigit(c) {
			if digitFlag {
				return false
			}
			digitFlag = true
		} else {
			digitFlag = false
		}
	}
	return true
}

func printString(s string) string {
	if len(s) == 0 {
		return ""
	}
	res := make([]rune, 0, len(s))
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsDigit(rune(s[i+1])) {
			count, err := strconv.Atoi(string(s[i+1]))
			if err != nil {
				log.Fatal(err.Error())
			}
			for j := 0; j < count; j++ {
				res = append(res, rune(s[i]))
			}
		} else if !unicode.IsDigit(rune(s[i])) {
			res = append(res, rune(s[i]))
		}
	}

	res = append(res, rune(s[len(s)-1]))
	return string(res)
}
