package observer_test

import (
	"oop/observer"
	"testing"
)

func TestObserver(t *testing.T) {
	sub := &observer.Subject{}
	sub.Add(&observer.Observer1{})
	sub.Add(&observer.Observer2{})
	sub.Add(&observer.Observer3{})
	sub.Notify("hi")
}
