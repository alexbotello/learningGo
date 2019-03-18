package main

import "fmt"

type human interface {
	speak(string) string
}

type engineer struct {
	name string
	age  int
}

func (e engineer) build() string {
	return fmt.Sprintf("%s is building cool new things", e.name)
}

func (e engineer) speak(phrase string) string {
	return fmt.Sprintf("%s has said: %s", e.name, phrase)
}

type teacher struct {
	name string
	age  int
}

func (t teacher) teach() string {
	return fmt.Sprintf("%s is now teaching students", t.name)
}

func (t teacher) speak(phrase string) string {
	return fmt.Sprintf("%s has said: %s", t.name, phrase)
}

func tellJoke(h human) string {
	return h.speak("How much would could a woodchuck chuck if a woodchuck could chuck would?")
}

func main() {
	billy := engineer{
		name: "Billy",
		age:  24,
	}

	susan := teacher{
		name: "Susan",
		age:  21,
	}

	fmt.Println(susan.teach())
	fmt.Println(billy.build())
	fmt.Println(tellJoke(billy))
}
