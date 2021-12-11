package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


// Flags структура с флагами
type Flags struct {
	ints  []*int
	bools []*bool
}
// GetFlags получение слайсов с флагами
func GetFlags() ([]*int, []*bool) {
	A := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	B := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	C := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "\"count\" (количество строк)")
	i := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	v := flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line num\", печатать номер строки")
	arr1 := []*int{A, B, C}
	arr2 := []*bool{c, i, v, F, n}
	return arr1, arr2
}

func openFile(path string) (map[int]*string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	mapStr := make(map[int]*string)
	sc := bufio.NewScanner(f)
	var i int = 1
	for sc.Scan() {
		b := sc.Text()
		mapStr[i] = &b
		i++
	}
	defer f.Close()
	return mapStr, nil
}

func getString(mapStr map[int]*string, word string) string {
	for _, value := range mapStr {
		if strings.Contains(*value, word) {
			return *value
		}
	}
	return ""
}

func grepBefore(mapStr map[int]*string, str string, num int) {
	for key, val := range mapStr {
		if *val == str {
			key = key - num
			fmt.Print("-----------------------------------------------------\n",*mapStr[key],"\n")
			for i := 0; i < num; i++ {
				key = key + 1
				fmt.Print(*mapStr[key],"\n")
			}
		}
	}
}

func grepAfter(mapStr map[int]*string, str string, num int) {
	for key, val := range mapStr {
		if *val == str {
			fmt.Print("-----------------------------------------------------\n")
			for i := 0; i <= num; i++ {
				key = key + i
				fmt.Print(*mapStr[key], "\n")
			}
		}
	}
}

func grepContext(mapStr map[int]*string, str string, num int){
	for key, val := range mapStr {
		if *val == str {
			key = key - num
			fmt.Print("-----------------------------------------------------\n",*mapStr[key],"\n")
			for i := num*(-1); i < num; i++{
				key = key + 1
				fmt.Print(*mapStr[key], "\n")
			}
		}
	}
}

func grepCount(mapStr map[int]*string, word string){
	var count int = 0
	for _, value := range mapStr {
		if strings.Contains(*value, word) {
			count++
		}
	}
	fmt.Printf("Количество строк, со словом %v: %v", word, count)
}

func grepIgnore(mapStr map[int]*string, word string) string{
	word = strings.ToLower(word)
	newmap := make(map[int]string)
	for key, val := range mapStr {
		newmap[key] = strings.ToLower(*val)
	}
	for key, value := range newmap {
		if strings.Contains(value, word) {
			return *mapStr[key]
		}
	}
	return ""
}

func grepInvert(mapStr map[int]*string, word string) {
	for _, val := range mapStr {
		if !strings.Contains(*val, word){
			fmt.Print(*val, "\n")
		}
	}
}

func grepFixed(mapStr map[int]*string, word string) *string{
	for _, value := range mapStr {
		if reflect.DeepEqual([]rune(word), []rune(*value)){
			return value
		}
	}
	return nil
}

func grepNum(mapStr map[int]*string, str string){
	for key, val := range mapStr{
		if *val == str{
			fmt.Printf("%v : %v", key, *val)
		}
	}
}

func main() {
	// Получение флагов
	flags := new(Flags)
	flags.ints, flags.bools = GetFlags()
	flag.Parse()
	var path string
	var word string
	if len(flag.Args()) != 0 {
		word = flag.Args()[0]
		path = flag.Args()[1]
	} else {
		fmt.Print("Вы не ввели имя файла")
		return
	}
	strMap, err := openFile(path)
	if err != nil {
		fmt.Println(err)
	}
	str := getString(strMap, word)
	if *flags.ints[0] != 0 {
		num := *flags.ints[0]
		if str == "" {
			fmt.Println("Строк не было найдено")
			return
		}
		grepAfter(strMap, str, num)
		return
	}
	if *flags.ints[1] != 0{
		if str == "" {
			fmt.Println("Строк не было найдено")
			return
		}
		num := *flags.ints[1]
		grepBefore(strMap, str, num)
		return
	}
	if *flags.ints[2] != 0{
		if str == "" {
			fmt.Println("Строк не было найдено")
			return
		}
		num := *flags.ints[2]
		grepContext(strMap, str, num)
		return
	}
	if *flags.bools[0] {
		grepCount(strMap, word)
		return
	}
	if *flags.bools[1] {
		str = grepIgnore(strMap, word)
		if str == ""{
			fmt.Print("Строка не найдена")
			return
		}
		fmt.Print(str)
		return
	}
	if *flags.bools[2]{
		grepInvert(strMap, word)
		return
	}
	if *flags.bools[3]{
		s := grepFixed(strMap, word)
		if s == nil{
			fmt.Println("Строк не было найдено")
			return
		}
		fmt.Println(*s)
		return
	}
	if *flags.bools[4]{
		if str == "" {
			fmt.Println("Строк не было найдено")
			return
		}
		grepNum(strMap, str)
		return
	}
	if str == "" {
		fmt.Println("Строк не было найдено")
		return
	}
	fmt.Println(str)
}
