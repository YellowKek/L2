package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type flags struct {
	k *int
	n *bool
	r *bool
	u *bool
}

func (f *flags) getActiveFlag() string {
	if *f.k != 0 {
		return strconv.Itoa(*f.k)
	} else if *f.n {
		return "n"
	} else if *f.r {
		return "r"
	}
	return ""
}

func main() {
	var f flags
	f.k = flag.Int("k", 1, "указание колонки для сортировки")
	f.n = flag.Bool("n", false, "сортировать по числовому значению")
	f.r = flag.Bool("r", false, "сортировать в обратном порядке")
	f.u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	var fileName string
	fileName = flag.Arg(0)
	data := GetData(fileName)
	sort(data, f)
	if *f.r {
		slices.Reverse(data)
	}

	printSlice(data, *f.u)
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

func printSlice(s []string, u bool) {
	var temp string
	for i := 0; i < len(s); i++ {
		if u {
			if i == 0 {
				temp = s[i]
				fmt.Println(s[i])
			} else if s[i] != temp {
				fmt.Println(s[i])
			}
			temp = s[i]
		} else {
			fmt.Println(s[i])
		}
	}
}

func sort(data []string, f flags) {
	target := *f.k - 1
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			curString := strings.Split(data[i], " ")
			nextString := strings.Split(data[j], " ")

			if compareStrings(curString, nextString, target) {
				swap(data, i, j)
			}
		}
	}
}

func sortN(data []string) {
	var (
		cur      string
		next     string
		curNums  []byte
		nextNums []byte
	)
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			cur = data[i]
			next = data[j]
			curNums = getDigits(cur)
			nextNums = getDigits(next)
			if len(nextNums) != 0 && len(curNums) == 0 {
				swap(data, i, j)
			} else if len(nextNums) != 0 && len(curNums) != 0 {
				for k := range curNums {
					if nextNums[k] < curNums[k] {
						swap(data, i, j)
						break
					}
				}
			}
		}
	}
}

func swap(s []string, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// нужно ли свапать
func compareStrings(s1, s2 []string, target int) bool {
	for i := target; i >= 0; i-- {
		if len(s1)-1 >= i && len(s2)-1 < i {
			return true
		}
		if len(s2)-1 >= i && len(s1)-1 < i {
			return false
		}
		if (len(s1)-1 >= i && len(s2)-1 >= i) || (len(s1)-1 < target && len(s2)-1 < target) {
			if strings.Compare(strings.Join(s1, " "), strings.Join(s2, " ")) == 1 {
				return true
			}
		}
	}
	return false
}

// s2 < s1
func stringComparator(s1, s2 string) bool {
	m := min(len(s1), len(s2))
	for i := 0; i < m; i++ {
		if s2[i] < s1[i] {
			return true
		}
	}
	return false
}

func getDigits(s string) (result []byte) {
	result = make([]byte, 0, len(s))
	for _, i := range s {
		_, err := strconv.Atoi(string(i))
		if err == nil {
			result = append(result, byte(i))
		}
	}
	return
}
