package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Акпят", "апмрирыи"}
	result := getAnagramsSet(input)
	validateAnagramsSet(result)
	fmt.Println(result)
}

func getAnagramsSet(input []string) map[string][]string {
	result := make(map[string][]string)

	result[input[0]] = []string{input[0]}
	for _, word := range input {
		word = strings.ToLower(word)
		_, ok := result[word]
		if !ok {
			flag := false
			for key := range result {
				if compareWords(getCharsMap(key), getCharsMap(word)) {
					flag = true
					result[key] = append(result[key], word)
					break
				}
			}
			if !flag {
				result[word] = append(result[word], word)
			}
		}
	}

	return result
}

func compareWords(w1, w2 map[rune]int) bool {
	if len(w1) != len(w2) {
		return false
	}
	for key, value := range w1 {
		value2, ok := w2[key]
		if !ok {
			return false
		}
		if value != value2 {
			return false
		}
	}
	return true
}

func getCharsMap(s string) map[rune]int {
	res := make(map[rune]int)
	for _, c := range s {
		v, ok := res[c]
		if !ok {
			res[c] = 1
		} else {
			res[c] = v + 1
		}
	}
	return res
}

func validateAnagramsSet(m map[string][]string) {
	for key, value := range m {
		if len(value) < 2 {
			delete(m, key)
		} else {
			slices.SortFunc(m[key], func(a, b string) int {
				return cmp.Compare(a, b)
			})
		}
	}
}
