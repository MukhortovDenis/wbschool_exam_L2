package main

import (
	"testing"
)

func Test_GetFile(t *testing.T) {
	t.Run("Test of getting file", func(t *testing.T) {
		_, err := GetFile("text.txt")
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_GetStrings(t *testing.T) {
	t.Run("Test of getting strings", func(t *testing.T) {
		buf, err := GetFile("text.txt")
		if err != nil {
			t.Error(err)
		}
		_, err = GetStrings(buf)
		if err != nil {
			t.Error(err)
		}

	})
}

func Test_CreateFile(t *testing.T) {
	t.Run("Test of creating file", func(t *testing.T){
		path := "test.txt"
		text := "testing testing testing"
		err := CreateFile(path, text)
		if err != nil {
			t.Error(err)
		}
	})

}

func Test_defaultSort(t *testing.T){
	t.Run("Testing default sort", func(t *testing.T){
		text:= "b c d a"
		expectedText := "a b c d"
		text = defaultSort(text)
		if text != expectedText{
			t.Error("Стандартная сортировка не сработала корректно")
		}
	})
}

func Test_SortK(t *testing.T){
	t.Run("Testing sorting with k-flag", func(t *testing.T){
		text:= "b c d a"
		expectedText := "d a b c"
		text = SortK(text, 2)
		if text != expectedText{
			t.Error("Сортировка с указанием колонки не сработала корректно")
		}
	})
}

func Test_SortB(t *testing.T){
	t.Run("Testing sorting with b-flag", func(t *testing.T){
		text:= "b c d a"
		expectedText := "bcda"
		text = SortB(text)
		if text != expectedText{
			t.Error("Игноривание пробелов не успешно")
		}
	})
}

func Test_SortM(t *testing.T){
	t.Run("Testing sorting with M-flag", func(t *testing.T){
		text:= "february november 123october june"
		expectedText := "february june 123october november"
		text = SortM(text)
		if text != expectedText{
			t.Error("Сортировка с месяцами не сработала корректно")
		}
	})
}

func Test_SortU(t *testing.T){
	t.Run("Testing sorting with u-flag", func(t *testing.T){
		text:= "b c a a"
		expectedText := "b c a"
		text = SortU(text)
		if text != expectedText{
			t.Error("Игноривание повторяющихся колонок завершилось неправильно")
		}
	})
}

func Test_SortR(t *testing.T){
	t.Run("Testing sorting with r-flag", func(t *testing.T){
		text:= "b c a d"
		expectedText := "d a c b"
		text = SortR(text)
		if text != expectedText{
			t.Error("Инверсия завершилась неправильно")
		}
	})
}

func Test_SortN(t *testing.T){
	t.Run("Testing sorting with n-flag", func(t *testing.T){
		text:= "14b c3 1a d45"
		expectedText := "1a c3 14b d45"
		text, err := SortN(text)
		if err != nil{
			t.Error(err)
		}
		if text != expectedText{
			t.Error("Сортировка по значениям завершилась неправильно")
		}
	})
}

func Test_SortC(t *testing.T){
	t.Run("Testing sorting with c-flag", func(t *testing.T){
		text:= "14b c3 1a d45"
		tmp := []string{"1a, c3, 14b, d45"}
		val := SortC(text, tmp)
		if val{
			t.Error("Исходный массив равен массиву новому")
		}
	})
}