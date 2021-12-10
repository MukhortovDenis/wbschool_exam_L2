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
	arr [6]*bool
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

func GetFlags() (*int, [6]*bool) {
	k := flag.Int("k", 0, "Сортировать по колонке")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	M := flag.Bool("M", false, "Сортировать по месяцам(англ)")
	b := flag.Bool("b", false, "Игнорировать пробелы")
	c := flag.Bool("c", false, "Проверка")
	arr := [6]*bool{n, r, u, M, b, c}
	return k, arr
}

func CreateFile(path string, str string) error {
	f, err := os.Create("new_" + path)
	if err != nil {
		return err
	}
	s := strings.Split(str, " ")
	for _, elem := range s {
		_, err := f.WriteString(elem + " \n")
		if err != nil {
			return err
		}
	}
	defer f.Close()
	return nil
}

func defaultSort(str string) string {
	s := strings.Split(str, " ")
	sort.Strings(s)
	return strings.Join(s, " ")
}

func SortK(str string, k int) string {
	s := strings.Split(str, " ")
	s[k], s[0] = s[0], s[k]
	sort.Strings(s[1:])
	return strings.Join(s, " ")
}

func SortB(str string) string {
	s := strings.Split(str, " ")
	return strings.Join(s, "")
}

func SortM(str string) string {
	month := [12]string{"january", "february", "march", "april", "august",
		"may", "june", "july", "september", "october", "november", "december"}
	s := strings.Split(str, " ")
	newS := []string{}
	for _, elemMonth := range month {
		for i, elem := range s {
			if strings.Contains(elem, elemMonth) {
				newS = append(newS, elem)
				copy(s[i:], s[i+1:])
				s[len(s)-1] = ""
				s = s[:len(s)-1]
			}
		}

	}
	for _, elem := range s {
		newS = append(newS, elem)
	}
	str = strings.Join(newS, " ")
	return str
}

func SortU(str string) string {
	s := strings.Split(str, " ")
	newS := []string{}
	keys := make(map[string]bool)
	for _, i := range s {
		if value := keys[i]; !value {
			keys[i] = true
			newS = append(newS, i)
		}
	}
	str = strings.Join(newS, " ")
	return str
}

func SortR(str string) string {
	s := strings.Split(str, " ")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	str = strings.Join(s, " ")
	return str
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
			if err != nil {
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
	strn := []string{}
	for _, num := range nums {
		for i, elem := range table {
			if num == elem {
				strn = append(strn, i)
			}
		}
	}
	ints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, elem := range s {
		var check bool = false
		for _, elemInt := range ints {
			if strings.Contains(elem, elemInt) {
				check = true
				break
			}
		}
		if !check {
			strn = append(strn, elem)
		}
	}
	return strings.Join(strn, " "), nil
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
	if *flags.k != 0 {
		str = SortK(str, *flags.k)
	}
	if *flags.arr[0] {
		str, err = SortN(str)
		if err != nil {
			fmt.Println("Ошибка сортировки чисел")
			os.Exit(0)
		}
	}
	if *flags.arr[1] {
		str = SortR(str)
	}
	if *flags.arr[2] {
		str = SortU(str)
	}
	if *flags.arr[3] {
		str = SortM(str)
	}
	if *flags.arr[4] {
		str = SortB(str)
	}
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
	err = CreateFile(path, str)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
