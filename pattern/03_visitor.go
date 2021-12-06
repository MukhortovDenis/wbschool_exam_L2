package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Visitor позволяет добавлять новые операции без изменения структур
// Так же можно применить одну операцию к разным обьектам
// Легко добавлять новые операции для структур, но если сами структуры меняются или их вложенность, то паттерн не использовать
type Visitor interface {
	visitFirst(first FirstElem) string
	visitSecond(second SecondElem) string
	visitThird(third ThirdElem) string
	visitObject(obj Object) string
}

type Element interface {
	Accept(vis Visitor) string
}
type FirstElem struct {
	msg string
}
type SecondElem struct {
	msg string
}
type ThirdElem struct {
	msg string
}
type Object struct {
	first  FirstElem
	second SecondElem
	third  ThirdElem
}

type PrintVisitor struct{}

func (pv *PrintVisitor) visitFirst(first FirstElem) string {
	return fmt.Sprintln("Посетил", first.msg)
}

func (pv *PrintVisitor) visitSecond(second SecondElem) string {
	return fmt.Sprintln("Посетил", second.msg)
}

func (pv *PrintVisitor) visitThird(third ThirdElem) string {
	return fmt.Sprintln("Посетил", third.msg)
}

func (pv *PrintVisitor) visitObject(obj Object) string {
	return fmt.Sprintf("Посетил object\n")
}

func (e *FirstElem) Accept(vis Visitor) string {
	return vis.visitFirst(*e)
}

func (e *SecondElem) Accept(vis Visitor) string {
	return vis.visitSecond(*e)
}

func (e *ThirdElem) Accept(vis Visitor) string {
	return vis.visitThird(*e)
}

func (obj *Object) Accept(vis Visitor) string {
	elements := []Element{
		&obj.first,
		&obj.second,
		&obj.third,
	}
	result := vis.visitObject(*obj)
	for _, elem := range elements {
		result += elem.Accept(vis)
	}
	return result
}

func main(){
	object := Object{
		first: FirstElem{
			msg: "Первый элемент",
		},
		second: SecondElem{
			msg: "Второй элемент",
		},
		third: ThirdElem{
			msg: "Третий элемент",
		},
	}
	vis := PrintVisitor{}
	result := object.Accept(&vis)
	fmt.Println(result)
}