package observer

// 抽象目标，被观察者接口
type ISubject interface {
	Register(observer IObserver, callback func(arg any))
	Unregister(observer IObserver)
	Notify(msg any)
}

// 观察者接口
type IObserver interface{}

// 具体目标
type EventBus struct {
	// 存放所有的观察者
	observers map[IObserver]func(arg any)
}

// 注册观察者
func (e *EventBus) Register(observer IObserver, callback func(arg any)) {
	e.observers[observer] = callback
}

// 取消注册观察者
func (e *EventBus) Unregister(observer IObserver) {
	delete(e.observers, observer)
}

// 通知观察者
func (e *EventBus) Notify(arg any) {
	for _, callback := range e.observers {
		callback(arg)
	}
}

// 具体的观察者
type Observer1 struct{}
type Observer2 struct{}
