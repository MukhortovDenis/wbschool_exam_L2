package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Хорош в разных представлениях обьекта, создавая его пошагово, но минус в том,
// что он достаточно большой и значительно усложняет код программы

type Builder struct {
	firstElement  string
	secondElement string
	thirdElement  string
}

type anotherBuilder struct {
	firstElement  string
	secondElement string
	thirdElement  string
}

type Launch struct {
	firstElement  string
	secondElement string
	thirdElement  string
}

type BuilderInterface interface {
	setFirst()
	setSecond()
	setThird()
	launch() Launch
}

type Director struct {
	bilder BuilderInterface
}

func NewBuilder() *Builder {
	return &Builder{}
}
func NewAnotherBuilder() *anotherBuilder {
	return &anotherBuilder{}
}

func (b *Builder) setFirst() {
	b.firstElement = "Первый элемент"
}
func (b *Builder) setSecond() {
	b.secondElement = "Второй элемент"
}
func (b *Builder) setThird() {
	b.thirdElement = "Третий элемент"
}

func (b *anotherBuilder) setFirst() {
	b.firstElement = "Первый элемент другого билдера"
}
func (b *anotherBuilder) setSecond() {
	b.secondElement = "Второй элемент другого билдера"
}
func (b *anotherBuilder) setThird() {
	b.thirdElement = "Третий элемент другого билдера"
}

func (b Builder) launch() Launch {
	return Launch{
		firstElement:  b.firstElement,
		secondElement: b.secondElement,
		thirdElement:  b.thirdElement,
	}
}

func (b anotherBuilder) launch() Launch {
	return Launch{
		firstElement:  b.firstElement,
		secondElement: b.secondElement,
		thirdElement:  b.thirdElement,
	}
}

func newDirector(b BuilderInterface) *Director {
	return &Director{
		bilder: b,
	}
}

func (d *Director) setBuilder(b BuilderInterface) {
	d.bilder = b
}
func (d Director) buildToLaunch() Launch {
	d.bilder.setFirst()
	d.bilder.setSecond()
	d.bilder.setThird()
	return d.bilder.launch()
}

func main() {
	builder := NewBuilder()
	director := newDirector(builder)
	launch := director.buildToLaunch()
	fmt.Println(launch.firstElement)
	//-----------------------------------------------------------------------------
	anotherBuilder := NewAnotherBuilder()
	director.setBuilder(anotherBuilder)
	launch = director.buildToLaunch()
	fmt.Println(launch.firstElement)
}
