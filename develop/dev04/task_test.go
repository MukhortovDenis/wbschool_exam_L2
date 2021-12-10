package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_searchAnagramma(t *testing.T) {
	checkMap := make(map[string][]string)
	a1 := []string{"кот", "кто", "окт", "ток"}
	a2 := []string{"кит", "тик"}
	a3 := []string{"крона", "норка"}
	a4 := []string{"пятка", "тяпка"}
	checkMap["ток"] = a1
	checkMap["тик"] = a2
	checkMap["норка"] = a3
	checkMap["тяпка"] = a4

	t.Run("testAnagram", func(t *testing.T) {
		massiv := &[]string{"Тяпка", "пятка", "ток", "коТ", "окт", "тик", "кИт", "Яблоко", "КТО", "НОРКа", "КроНА"}
		getMap := searchAnagramma(massiv)
		fmt.Println(checkMap)
		for key, value := range *getMap {
			if !reflect.DeepEqual(value, checkMap[key]) {
				t.Errorf("anagramma() = %v, want %v", getMap, checkMap)
			}
		}

	})
}
