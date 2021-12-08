package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
// Как я понял в Go можно роеализовать только простую фабрику
// Выделяет код производства продуктов в одно место, упрощая поддержку кода
type Factory interface {
	setName(name string)
	setInfo(info string)
	getName() string
	getInfo() string
}

type Object struct {
	name string
	info string
}

type ConcreteObject1 struct {
	Object
}

type ConcreteObject2 struct {
	Object
}

func (obj *Object) setName(name string) {
	obj.name = name
}

func (obj *Object) setInfo(info string) {
	obj.info = info
}

func (obj *Object) getName() string {
	return obj.name
}

func (obj *Object) getInfo() string {
	return obj.info
}

func newConcreteObject1() Factory {
	return &ConcreteObject1{
		Object: Object{
			name: "Первый обьект",
			info: "Cоздали первый обьект",
		},
	}
}

func newConcreteObject2() Factory {
	return &ConcreteObject2{
		Object: Object{
			name: "Второй обьект",
			info: "Cоздали второй обьект",
		},
	}
}

func getObj(number int) Factory {
	if number == 1 {
		return newConcreteObject1()
	}
	if number == 2 {
		return newConcreteObject2()
	}
	return nil
}

func main() {
	obj1 := getObj(1)
	obj2 := getObj(2)
	fmt.Println(obj1.getName())
	obj1.setInfo("Обновили информацию о первом")
	fmt.Println(obj1.getInfo())
	fmt.Println(obj2.getInfo())
	fmt.Println(obj2.getName())
}
