package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	f bool
	n bool
}

func main() {
	var f flags
	f.A = *flag.Int("A", 0, "печатать +N строк после совпадения")
	f.B = *flag.Int("B", 0, "печатать +N строк до совпадения")
	f.C = *flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	f.c = *flag.Bool("c", false, "количество строк")
	f.i = *flag.Bool("i", false, "игнорировать регистр")
	f.v = *flag.Bool("v", false, "вместо совпадения, исключать")
	f.f = *flag.Bool("f", false, "точное совпадение со строкой, не паттерн")
	f.n = *flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	var file, pattern string
	pattern = flag.Arg(0)
	file = flag.Arg(1)
	data := GetData(file)
	processing(pattern, data, f)
}

func processing(pattern string, data []string, f flags) {
	for _, str := range data {
		matched, err := regexp.MatchString(pattern, str)
		if err != nil {
			log.Fatal("IN MATCH: ", err.Error())
		}
		if matched {
			fmt.Println(str)
		}
	}
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
