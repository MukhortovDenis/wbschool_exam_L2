package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func GetRune(s interface{}) ([]rune, error) {
	var a interface{} = s
	if str, isStr := a.(string); isStr {
		if str == "" {
			return nil, errors.New("результат распаковки \"\"(пустая строка)")
		}
		return []rune(str), nil
	}
	if i, isInt := a.(int); isInt {
		return nil, fmt.Errorf("вы ввели число %v", i)
	}
	return nil, errors.New("неизвестный формат")
}

func CheckRune(r []rune) ([]rune, error) {
	if r[0] >= 48 && r[0] <= 57 {
		return nil, fmt.Errorf("вы ввели некорректную строку, которая начинается с %v", string(r[0]))
	}
	if r[0] == 32 {
		return nil, errors.New("ваша строка начинается с пробела")
	}
	if r[0] == 92 {
		return nil, errors.New("ваша строка начинается с обратного слэша")
	}
	if r[len(r)-1] == 92{
		return nil, errors.New("ваша строка заканчивается обратным слэшем")
	}
	return r, nil
}

func Unpacking(r []rune) (string, error) {
	str := new(strings.Builder)
	for i := 0; i < len(r); i++ {
		switch {	
			case r[i] == 92 && r[i+1] >= 48 && r[i+1] <= 57:
				_, err := str.WriteString(string(r[i+1]))
				if err != nil {
					return "", err
				}
				i++
			case r[i] == 92 && r[i+1] == 92:
				_, err := str.WriteString(string(r[i+1]))
				if err != nil {
					return "", err
				}
				i++
			case r[i] >= 48 && r[i] <= 57:
				num, err := strconv.Atoi(string(r[i]))
				if err != nil {
					return "", err
				}
				for ; num > 1; num-- {
					_, err := str.WriteString(string(r[i-1]))
					if err != nil {
						return "", err
					}
				}
			default:
				_, err := str.WriteString(string(r[i]))
				if err != nil {
					return "", err
				}
			}
	}

	return str.String(), nil
}

func main() {
	str, err := GetRune(`qwe\\5`)
	if err != nil {
		fmt.Println(err)
		return
	}
	str, err = CheckRune(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	unpack, err := Unpacking(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(unpack)
}
