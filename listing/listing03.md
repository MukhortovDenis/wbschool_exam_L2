Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Интерфейсы реализованы в виде двух элементов, тип T и значение V. V значение никогда не является самим интерфейсом
и имеет тип T. Даже когда мы задаем nil это уже имеет тип(в нашем случае *os.PathError). Поэтому проверка на нил будет false, т.к
nil справа ознает пустой интерфейс, а слева интерфейс с типом *os.PathError, в котором V = nil

```
