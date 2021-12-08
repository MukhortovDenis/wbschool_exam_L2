package pattern
import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
// Этот паттерн позволяет выбирать алгоритма в ходе исполнения, все алгоритма взаимозаменяемы
// Но из-за большого количества структур усложняет код и пользоваться можно только при понимании подходящей статегии
type Eviction interface {
	strategyResult(obj *Object)
}

type Strategy1 struct{}
type Strategy2 struct{}
type Strategy3 struct{}

type Object struct {
	Info     string
	Strategy Eviction
}

func (obj *Object) initStrategy(info string, strategy Eviction) {
	obj.Info = info
	obj.Strategy = strategy
}

func (obj *Object) changeStrategy(strategy Eviction) {
	obj.Strategy = strategy
}

func (str *Strategy1) strategyResult(obj *Object) {
	fmt.Println(obj.Info)
}

func (str *Strategy2) strategyResult(obj *Object) {
	fmt.Println(obj.Info + obj.Info)
}

func (str *Strategy3) strategyResult(obj *Object) {
	fmt.Println(obj.Info + obj.Info + obj.Info)
}

func (obj *Object) showInfo() {
	obj.Strategy.strategyResult(obj)
}

func main() {
	obj := &Object{}
	str1 := &Strategy1{}
	str2 := &Strategy2{}
	str3 := &Strategy3{}
	obj.initStrategy("something", str1)
	obj.showInfo()
	obj.changeStrategy(str2)
	obj.showInfo()
	obj.changeStrategy(str3)
	obj.showInfo()
}
