package main

import "fmt"

type Writer interface {
	Write([]byte) error
}

type Singer interface {
	Sing()
}

type Bird struct{}

func (b Bird) Sing() {
	fmt.Println("汪汪汪")
}

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("猫")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("狗")
}

func main() {
	// c := Cat{}
	// c.Say()
	// d := Dog{}
	// d.Say()
	// s := Sheep{}
	// s.Say()

	var c Cat
	MakeHungry(c)
	var d Dog
	MakeHungry(d)
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("羊")
}

func MakeCatHungry(c Cat) {
	c.Say()
}
func MakeSheepHungry(s Sheep) {
	s.Say()
}

type Sayer interface {
	Say()
}

func MakeCatHungry1(s Sayer) {
	s.Say()
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}
