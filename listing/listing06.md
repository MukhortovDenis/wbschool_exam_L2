Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Мы сможем изменить только первый элемент так, как в функции modifySlice в начале мы работаем с одним массивом
Но когда используем append, создается новый массив и указатель i теперь принадлежит ему
Потом создается еще один массив уже из пяти элементов
Поскольку в функции мы не возвращаем указатель, то после modifySlice мы работает с самым первым массивом

```
