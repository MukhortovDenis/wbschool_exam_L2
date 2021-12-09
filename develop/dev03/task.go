package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type flagStruct struct {
	k   *int
	arr [7]*bool
}

func GetFile(filename string) (*bytes.Buffer, error) {
	text, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	buffer := bytes.Buffer{}
	sc := bufio.NewScanner(text)
	for sc.Scan() {
		buffer.WriteString(sc.Text())
	}
	defer text.Close()
	return &buffer, nil
}

func GetStrings(buffer *bytes.Buffer) (string, error) {
	sb := new(strings.Builder)
	for buffer.Len() > 0 {
		_, err := sb.WriteString(string(buffer.Next(1)))
		if err != nil {
			return "", err
		}
	}
	return sb.String(), nil
}

func GetFlags() (*int, [7]*bool) {
	k := flag.Int("k", 0, "Сортировать по колонке")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	M := flag.Bool("M", false, "Сортировать по месяцам(англ)")
	b := flag.Bool("b", false, "Игнорировать пробелы")
	c := flag.Bool("c", false, "Проверка")
	h := flag.Bool("h", false, "Сортировать по числовому значению с учетом суффиксов")
	arr := [7]*bool{n, r, u, M, b, c, h}
	return k, arr
}

func defaultSort(str string) string {
	s := strings.Split(str, "\n")
	fmt.Println(s[0])
	sort.Strings(s)
	return strings.Join(s, " ")
}

func SortN(str string) (string, error) {
	s := strings.Split(str, " ")
	table := make(map[string]int)
	for _, elem := range s {
		b := []rune(elem)
		d := []rune{}
		for i := range b {
			if b[i] >= 48 && b[i] <= 57 {
				d = append(d, b[i])
			}
		}
		newstr := string(d)
		if newstr != "" {
			num, err := strconv.Atoi(newstr)
			if err != nil{
				return "", err
			}
		table[elem] = num
		}
	}
	nums := []int{}
	for _, num := range table {
		nums = append(nums, num)
	}
	sort.Ints(nums)
	s = []string{}
	for _, num := range nums{
		for i, elem := range table{
			if num == elem{
				s = append(s, i)
			}
		}
	}
	return strings.Join(s, " "), nil
}

func SortC(str string, tmp []string) bool {
	strOld := strings.Join(tmp, " ")
	if strOld == str {
		return false
	}
	return true
}

func main() {
	flags := new(flagStruct)
	flags.k, flags.arr = GetFlags()
	flag.Parse()
	var path string
	if len(flag.Args()) != 0 {
		path = flag.Args()[0]
	} else {
		fmt.Print("Вы не ввели имя файла")
		return
	}
	newbuf, err := GetFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str, err := GetStrings(newbuf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(str, " ")
	tmp := make([]string, len(s))
	copy(tmp, s)
	str = defaultSort(str)
	if *flags.arr[5] {
		ok := SortC(str, tmp)
		if !ok {
			fmt.Println("Данные не отсортированы")
			os.Exit(0)
		} else {
			fmt.Println("Данные отсортированы")
			os.Exit(0)
		}
	}
	if *flags.arr[0] {
		str, err = SortN(str)
		if err != nil{
			fmt.Println("Ошибка сортировки чисел")
			os.Exit(0)
		}
	}
	if *flags.arr[1] {

	}
	if *flags.arr[2] {

	}
	if *flags.arr[3] {

	}
	if *flags.arr[4] {

	}
	if *flags.k != 0 {

	}
	fmt.Println(str)
}
