Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	err := new(customError)
	return err
}

func main() {
	err := test()
	if err.Error() != "" {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
В соответствии с listing03 получаем nil типа, а не просто nil
В самой функции test, возвращается nil(поэтому err остается типом error), но нам нужен тип *customError, получаю указатель на структуру
и возвращаю его.
Теперь к err сможем применить метод Error, который возвращает пустую строку, если в msg ничего нет
Делаем if на проверку нуля в строке(пустая строка "") 

```
