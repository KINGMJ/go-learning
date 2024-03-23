package observer

import "fmt"

// 抽象目标，被观察者接口
type ISubject interface {
	Add(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// 观察者接口
type IObserver interface {
	Update(msg string) // 观察到别人发生变化，自己也改变
}

// 具体目标
type Subject struct {
	// 存放所有的观察者
	observers []IObserver
}

func (sub *Subject) Add(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, ob := range sub.observers {
		ob.Update(msg)
	}
}

// 具体的观察者
type Observer1 struct{}

func (Observer1) Update(msg string) {
	fmt.Printf("Observer1 响应信息: %s\n", msg)
}

type Observer2 struct{}

func (Observer2) Update(msg string) {
	fmt.Printf("Observer2 响应信息：%s\n", msg)
}

type Observer3 struct{}

func (Observer3) Update(msg string) {
	fmt.Printf("Observer3 响应信息：%s\n", msg)
}
