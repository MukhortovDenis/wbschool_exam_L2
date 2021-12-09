package main

import (
	"errors"
	"log"
	"testing"
	"fmt"
)


func Test_getRune(t *testing.T) {
	type TestCase struct {
		Data   interface{}
		Result string
		Error  error
	}
	cases := []TestCase{
		TestCase{"a4bc2d5e", "a4bc2d5e", nil},
		TestCase{"", "", errors.New("результат распаковки \"\"(пустая строка)")},
		TestCase{`qwe\4\5`, `qwe\4\5`, nil},
		TestCase{45, "45", fmt.Errorf("вы ввели число")},
	}
	for caseNum, item := range cases {
		_, err := GetRune(item.Data)
		if err == item.Error {
			log.Printf("кейс %v прошел проверку", caseNum)
		}
		if err != nil && item.Error != nil {
			log.Printf("кейс %v прошел проверку", caseNum)
		}
		if err != nil && item.Error == nil {
			t.Errorf("кейс %v не прошел проверку", caseNum)
		}
	}
}

func Test_CheckRune(t *testing.T){
	type TestCase struct {
		Data   []rune
		Result []rune
		Error  error
	}
	cases := []TestCase{
		TestCase{[]rune("4bc2d5e"), nil , fmt.Errorf("вы ввели некорректную строку, которая начинается c числа")},
		TestCase{[]rune(" a4bc2d5e"), nil, errors.New("начали с новой строки")},
		TestCase{[]rune(`\выаывп`), nil, errors.New("начали с слеша")},
		TestCase{[]rune(`fg4j\`), nil, errors.New("закончили слешем")},
		TestCase{[]rune(`qwe\\5`), []rune(`qwe\\5`), nil},
	}
	for caseNum, item := range cases {
		_, err := CheckRune(item.Data)
		if err == item.Error {
			log.Printf("кейс %v прошел проверку", caseNum)
		}
		if err != nil && item.Error != nil {
			log.Printf("кейс %v прошел проверку", caseNum)
		}
		if err != nil && item.Error == nil {
			t.Errorf("кейс %v не прошел проверку", caseNum)
		}
	}
}

func Test_Unpacking(t *testing.T){
	type TestCase struct {
		Data   []rune
		Result string
		Error  error
	}
	cases := []TestCase{
		TestCase{[]rune("a4bc2d5e"), "aaaabccddddde" , nil},
		TestCase{[]rune("abcd"), "abcd", nil},
		TestCase{[]rune(`qwe\4\5`), `qwe45`, nil},
		TestCase{[]rune(`qwe\45`), `qwe44444`, nil},
		TestCase{[]rune(`qwe\\5`), `qwe\\\\\`, nil},
	}
	
	for caseNum, item := range cases {
		result, err := Unpacking(item.Data)
		if err == item.Error && result == item.Result {
			log.Printf("кейс %v прошел проверку", caseNum)
		}
		if err != nil && item.Error == nil || result != item.Result {
			t.Errorf("кейс %v не прошел проверку", caseNum)
		}
	}
}