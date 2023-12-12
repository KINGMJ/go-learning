package main

import "fmt"

type Object struct {
	data       int
	references int
}

func NewObject(data int) *Object {
	return &Object{data: data, references: 1}
}

func (o *Object) AddReference() {
	o.references++
}

func (o *Object) ReleaseReference() {
	o.references--
	if o.references == 0 {
		fmt.Println("Object released:", o.data)
	}
}

func processObject(obj *Object) {
	fmt.Println("Processing object with data:", obj.data)
}

func main() {
	obj := NewObject(43)
	// obj.AddReference()
	processObject(obj)
	obj.ReleaseReference()
}
