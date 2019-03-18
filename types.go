package main

import "fmt"

// Human is the interface that wraps the basic Speak method
type Human interface {
	Speak(string) string
}

// Engineer type that refers to a Human that's an engineer
type Engineer struct {
	name string
	age  int
}

func (e Engineer) Build() string {
	return fmt.Sprintf("%s is building cool new things", e.name)
}

func (e Engineer) Speak(phrase string) string {
	return fmt.Sprintf("%s has said: %s", e.name, phrase)
}

// Teacher type refers to a Human that's a teacher
type Teacher struct {
	name string
	age  int
}

func (t Teacher) Teach() string {
	return fmt.Sprintf("%s is now teaching students", t.name)
}

func (t Teacher) Speak(phrase string) string {
	return fmt.Sprintf("%s has said: %s", t.name, phrase)
}

// TellJoke accepts a Human interface and returns a joke phrase
func TellJoke(h Human) string {
	return h.Speak("How much would could a woodchuck chuck if a woodchuck could chuck would?")
}
