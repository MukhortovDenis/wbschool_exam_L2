package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// os.Stdin - это то, к чему файловый дескриптор 0 подключается операционной системой.
// В зависимости от типа файлового дескриптора чтение будет вести себя по-разному.
// Если вы вызываете программу в командной строке, оболочка подключает ее к своему управляющему терминалу, который обычно буферизуется строками,
// и чтение из него блокируется до тех пор, пока не будет записана новая строка.
// Если вы введете в него другую команду, она подключит стандартный вывод этой программы к вашему стандартному входу, и как только эта программа (в вашем примере echo) завершится,
//  чтение вернет io.EOF.

// Flags структура с флагами
type Flags struct {
	f *string
	d *string
	s *bool
}

// GetFlags получение флагов
func GetFlags() (*string, *string, *bool) {
	f := flag.String("f", "", "выбрать поля (колонки)")
	d := flag.String("d", "\t", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")
	return f, d, s
}

func getFromStdin() map[int]*string {
	strMap := make(map[int]*string)
	var i int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		strMap[i] = &txt
		i++
	}
	return strMap
}

func getColumn(flag *string, strMap map[int]*string) map[int][]string {
	anotherMap := make(map[int][]string)
	for key, val := range strMap {
		if strings.Contains(*val, *flag) {
			slice := strings.Split(*val, *flag)
			anotherMap[key] = slice
		} else if strings.Contains(*val, "    ") {
			slice := strings.Split(*val, "    ")
			anotherMap[key] = slice
		} else {
			slice := []string{*val}
			fmt.Println(slice)
			anotherMap[key] = slice
		}
	}
	return anotherMap
}

func getSeparated(strMap map[int]*string, flag *string) map[int]*string {
	for k, v := range strMap {
		if !strings.Contains(*v, *flag) {
			delete(strMap, k)
		}
	}
	return strMap
}

func getFields(flag *string, newmap map[int][]string) {
	if strings.Contains(*flag, "-") {
		slice := strings.Split(*flag, "-")
		if slice[0] == "" {
			for _, val := range newmap {
				if val != nil {
					sb := strings.Builder{}
					for i := range val {
						j, err := strconv.Atoi(slice[1])
						if err != nil {
							fmt.Print(err)
							return
						}
						if i <= j {
							sb.WriteString(" ")
							sb.WriteString(val[i])

						}
					}
					fmt.Println(sb.String())
				}
			}
			return
		} else if slice[1] == "" {
			for _, val := range newmap {
				if val != nil {
					sb := strings.Builder{}
					for i := range val {
						j, err := strconv.Atoi(slice[0])
						if err != nil {
							fmt.Print(err)
							return
						}
						if i >= j {
							sb.WriteString(val[i])
							sb.WriteString(" ")
						}
					}
					fmt.Println(sb.String())
				}
			}
			return
		} else {
			for _, val := range newmap {
				sb := strings.Builder{}
				if val != nil {
					for i := range val {
						j, err := strconv.Atoi(slice[0])
						if err != nil {
							fmt.Print(err)
							return
						}
						k, err := strconv.Atoi(slice[1])
						if err != nil {
							fmt.Print(err)
							return
						}
						if i >= j && i <= k {
							sb.WriteString(val[i])
							sb.WriteString(" ")
						}
					}
					fmt.Println(sb.String())
				}
			}
			return
		}
	} else if strings.Contains(*flag, ",") {
		slice := strings.Split(*flag, ",")
		for _, val := range newmap {
			if val != nil {
				sb := strings.Builder{}
				for _, v := range slice {
					vInt, err := strconv.Atoi(v)
					if err != nil {
						fmt.Print(err)
						return
					}
					sb.WriteString(val[vInt])
					sb.WriteString(" ")
				}
				fmt.Println(sb.String())
			}
		}
	} else {
		k, err := strconv.Atoi(*flag)
		if err != nil {
			fmt.Print(err)
			return
		}
		for _, val := range newmap {
			if val != nil {
				sb := strings.Builder{}
				sb.WriteString(val[k])
				fmt.Println(sb.String())
			}
		}
	}

}

func main() {
	flags := new(Flags)
	flags.f, flags.d, flags.s = GetFlags()
	flag.Parse()
	strMap := getFromStdin()
	if *flags.s {
		strMap = getSeparated(strMap, flags.d)
	}
	newMap := getColumn(flags.d, strMap)
	if *flags.f != "" {
		getFields(flags.f, newMap)
		return
	}
}
