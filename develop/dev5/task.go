package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type flags struct {
	A *int
	B *int
	C *int
	c *bool
	i *bool
	v *bool
	F *bool
	n *bool
}

func main() {
	var f flags
	f.A = flag.Int("A", 0, "печатать +N строк после совпадения")
	f.B = flag.Int("B", 0, "печатать +N строк до совпадения")
	f.C = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	f.c = flag.Bool("c", false, "количество строк")
	f.i = flag.Bool("i", false, "игнорировать регистр")
	f.v = flag.Bool("v", false, "вместо совпадения, исключать")
	f.F = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	f.n = flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	var file, pattern string
	pattern = flag.Arg(0)
	file = flag.Arg(1)
	data := GetData(file)
	processing(pattern, data, f)
}

func processing(pattern string, data []string, f flags) {
	count := 0
	if *f.C != 0 {
		*f.A = *f.C
		*f.B = *f.C
	}
	lastMatchedInd := 0
	for i, s := range data {
		str := s
		needToPrint := false
		if *f.i {
			str = strings.ToLower(s)
			pattern = strings.ToLower(pattern)
		}
		if *f.F {
			if strings.Compare(pattern, str) == 0 {
				needToPrint = true
			}
		} else {
			matched, err := regexp.MatchString(pattern, str)
			if err != nil {
				log.Fatal("IN MATCH: ", err.Error())
			}
			if matched && !*f.v {
				needToPrint = true
			} else if !matched && *f.v {
				needToPrint = true
			}
		}
		if needToPrint {
			var before, after []string
			lastMatchedInd = i
			if *f.B != 0 {
				before = getStringsBefore(data, pattern, i, *f.B)
				for i2 := range before {
					fmt.Println(before[i2])
				}
			}
			if *f.c {
				count++
			} else if *f.n {
				fmt.Println(i + 1)
			} else {
				fmt.Println(s)
				if *f.A != 0 && *f.B == 0 {
					after = getStringsAfter(data, pattern, i+1, *f.A)
					for i2 := range after {
						fmt.Println(after[i2])
					}
				}
			}
		}
	}
	if *f.c {
		fmt.Println(count)
	} else if *f.C != 0 {
		after := getStringsAfter(data, pattern, lastMatchedInd+1, *f.A)
		for i2 := range after {
			fmt.Println(after[i2])
		}
	}
}

func getStringsBefore(s []string, pattern string, ind, i int) []string {
	temp := s[ind-i : ind]
	for i2, s2 := range temp {
		matched, err := regexp.MatchString(pattern, s2)
		if err != nil {
			log.Fatal(err.Error())
		}
		if matched {
			return temp[i2+1:]
		}
	}
	return temp
}

func getStringsAfter(s []string, pattern string, ind, i int) []string {
	temp := s[ind : ind+i]
	for i2, s2 := range temp {
		matched, err := regexp.MatchString(pattern, s2)
		if err != nil {
			log.Fatal(err.Error())
		}
		if matched {
			return temp[:i2]
		}
	}
	return temp
}

func GetData(s string) []string {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal("Файл не найден!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
