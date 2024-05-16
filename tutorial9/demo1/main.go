package main

import (
	"fmt"

	evbus "github.com/asaskevich/EventBus"
)

func main() {
	bus := evbus.New()
	bus.Subscribe("main:11", func(msg Message) {
		var a A
		a.calculator(msg)
	})

	var a A
	bus.Subscribe("main:11", a.print)

	bus.Publish("main:11", Message{Name: "message1", Content: map[string]any{"name": "jack", "age": 12}})
}

type Message struct {
	Name    string
	Content map[string]any
}

type A struct{}

func (a A) calculator(msg Message) {
	fmt.Println(msg)
}

func (a A) print(msg Message) {
	fmt.Printf("name is : %s", msg.Name)
}
