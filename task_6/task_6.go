package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type flags struct {
	f *string
	d *string
	s *bool
}

func main() {
	f := flags{}
	f.f = flag.String("f", "", "выбрать поля (колонки)")
	f.d = flag.String("d", "", "использовать другой разделитель")
	f.s = flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	file := flag.Arg(0)
	data := GetData(file)
	cut(f, data)
	//a, b := fArgsParse(*f.f)
	//fmt.Println()
}

func cut(f flags, data []string) {
	columns, flag := fArgsParse(*f.f)
	if len(*f.f) == 0 {
		fmt.Println("cut: you must specify a list of fields")
		return
	} else if len(*f.d) == 0 {
		for _, str := range data {
			fmt.Println(str)
		}
	} else {
		for _, datum := range data {
			delimited := strings.Split(datum, *f.d)
			if *f.s && !strings.Contains(datum, *f.d) {
				continue
			}

			if flag {
				if len(delimited) <= columns[0] {
					fmt.Println(datum)
				} else {
					for i := columns[0]; i < len(delimited); i++ {
						fmt.Print(delimited[i])
						if i < len(delimited)-1 {
							fmt.Print(*f.d)
						}
					}
					fmt.Println()
				}
			} else {
				for i, column := range columns {
					fmt.Print(delimited[column])
					if i < len(columns)-1 {
						fmt.Print(*f.d)
					}
					fmt.Println()
				}
			}

		}

	}
}

func fArgsParse(args string) ([]int, bool) {
	result := make([]int, 0, 2)
	flag := false
	if len(args) == 2 {
		if string(args[0]) == "-" {
			num, err := strconv.Atoi(string(args[1]))
			if err != nil {
				log.Fatal("Неверный формат")
			}
			for i := 0; i < num; i++ {
				result = append(result, i)
			}
		} else {
			num, err := strconv.Atoi(string(args[0]))
			if err != nil {
				log.Fatal("Неверный формат")
			}
			result = append(result, num-1)
			flag = true
		}
	} else {
		temp := strings.Split(args, ",")
		for _, s := range temp {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal("Неверный формат")
			}
			result = append(result, num-1)
		}
	}

	return result, flag
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
