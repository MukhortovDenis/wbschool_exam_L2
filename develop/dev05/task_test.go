package main

import "testing"

var (
mystr = "my String is created"
mystr2 = "my Another sTring is created"
mystr3 = "my Another sTring is created again"
mapstr = map[int]*string{1:&mystr, 2:&mystr2, 3:&mystr3}
)

func Test_openFile(t *testing.T){
	t.Run("Открытие файла", func(t *testing.T){
		_, err := openFile("text.txt")
		if err != nil {
			t.Error(err)
		}
		_, err = openFile("wrong.txt")
		if err == nil{
			t.Error("Ошибка")
		}
	})
}

func Test_getString(t *testing.T){
	t.Run("Получение строки", func(t *testing.T){
		word := "String"
		str := getString(mapstr, word)
		if mystr != str{
			t.Error("Строка не найдена")
		}
	})
}

func Test_grepBefore(t *testing.T){
	t.Run("Чтение строк до совпадения", func(t *testing.T){
		grepBefore(mapstr, mystr2, 1)
	})
}

func Test_grepAfter(t *testing.T){
	t.Run("Чтение строк после совпадения", func(t *testing.T){
		grepAfter(mapstr, mystr, 1)
	})
}

func Test_grepContext(t *testing.T){
	t.Run("Чтение строк до и после совпадения", func(t *testing.T){
		grepContext(mapstr, mystr2, 1)
	})
}

func Test_grepCount(t *testing.T){
	t.Run("Количество строк по слову", func(t *testing.T){
		grepCount(mapstr, "string")
	})
}

func Test_grepIgnore(t *testing.T){
	t.Run("Игнорирование регистра", func(t *testing.T){
		str := grepIgnore(mapstr, "AGAiN")
		if str != "my Another sTring is created again"{
			t.Errorf("Cтрока %v не равна \"my Another sTring is created again\"", str)
		}
	})
}

func Test_grepInvert(t *testing.T){
	t.Run("Исключение", func(t *testing.T){
		grepInvert(mapstr, "string")
	})
}

func Test_grepFixed(t *testing.T){
	t.Run("Поиск полной строки", func(t *testing.T){
		a := grepFixed(mapstr, "my Another sTring is created")
		if *a != "my Another sTring is created"{
			t.Error("Не найдене строки")
		}
	})
}

func Test_grepNum(t *testing.T){
	t.Run("Исключение", func(t *testing.T){
		grepNum(mapstr, "my Another sTring is created")
	})
}