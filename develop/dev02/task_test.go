package main

import (
	"errors"
	"log"
	"testing"
	"fmt"
)

type TestCase struct {
	Data   interface{}
	Result string
	Error  error
}

func Test_getRune(t *testing.T) {
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
