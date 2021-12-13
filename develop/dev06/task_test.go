package main


import (
	"testing"
	"reflect"
)
var (
mystr = "my\tString\tis\tcreated"
mystr2 = "my\tAnother\tsTring\tis\tcreated"
mystr3 = "my\tAnother\tsTring\tis\tcreated\tagain"
mapstr = map[int]*string{1:&mystr, 2:&mystr2, 3:&mystr3}
)

func Test_getColumn(t *testing.T){
	t.Run("Разбитие на колонки", func(t *testing.T){
		sign := "\t"
		str := getColumn(&sign, mapstr)
		slice1 := []string{"my", "String", "is", "created "}
		slice2 := []string{"my", "Another","sTring", "is", "created "}
		slice3 := []string{"my", "Another", "sTring", "is", "created", "again "}
		expected := map[int][]string{
			0:slice1,
			2:slice2,
			3:slice3}
		for k, v := range str{
			if reflect.DeepEqual(v, expected[k]){
				t.Errorf("Получили %v, ожидали %v", v, expected[k])
			}
		}
	})
}

func Test_getSepated(t *testing.T){
	t.Run("Получение строки только с разделителем", func(t *testing.T){
		sign := "\t"
		str := getSeparated(mapstr, &sign)
		slice1 := []string{"my", "String", "is", "created "}
		slice2 := []string{"my", "Another","sTring", "is", "created "}
		slice3 := []string{"my", "Another", "sTring", "is", "created", "again "}
		expected := map[int][]string{
			0:slice1,
			2:slice2,
			3:slice3}
		for k, v := range str{
			if reflect.DeepEqual(v, expected[k]){
				t.Errorf("Получили %v, ожидали %v", v, expected[k])
			}
		}
	})
}

func Test_getFields(t *testing.T){
	t.Run("Получение строки только с разделителем", func(t *testing.T){
		slice1 := []string{"my", "String", "is", "created "}
		slice2 := []string{"my", "Another","sTring", "is", "created "}
		slice3 := []string{"my", "Another", "sTring", "is", "created", "again "}
		expected := map[int][]string{
			0:slice1,
			2:slice2,
			3:slice3}
		sign := "1-"
		getFields(&sign, expected)
		sign = "-1"
		getFields(&sign, expected)
		sign = "0-2"
		getFields(&sign, expected)
		sign = "0,2"
		getFields(&sign, expected)
		sign = "1"
		getFields(&sign, expected)
	})
}