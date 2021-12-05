package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Смысл фасада дать простой интерфейс для использование без подробностей как он устроен и работает
// Будет ли его реализация больше плюсом или минусом зависит от того, нужно ли пользователю тестирование, дополнение, изменинение в целом этого интерфейса
type Facade struct {
	first  *first
	second *second
	third  *third
}

type first struct {
	msg string
}
type second struct {
	msg string
}
type third struct {
	msg string
}

func newFacade() *Facade {
	facade := &Facade{
		first: newFirst("Первый сегмент существует"),
		second: newSecond("Второй сегмент существует"),
		third: newThird("Третий сегмент существует"),
	}
	fmt.Println("Создание фасада")
	return facade
}

func newFirst(msg string) *first {
	return &first{
		msg: msg,
	}
}

func (f *Facade) checkFirst(){
	fmt.Print(f.first.msg)
}

func newSecond(msg string) *second {
	return &second{
		msg: msg,
	}
}
func (f *Facade) checkSecond(){
	fmt.Print(f.second.msg)
}
func newThird(msg string) *third {
	return &third{
		msg: msg,
	}
}
func (f *Facade) checkThird(){
	fmt.Print(f.third.msg)
}

func main() {
	facade := newFacade()
	facade.checkSecond()
}
