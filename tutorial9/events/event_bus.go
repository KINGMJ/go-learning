package events

import evbus "github.com/KINGMJ/go-learning/tutorial9/eventbus"

var eventBus = evbus.New()

func Publish(event Event) error {
	return eventBus.Publish(string(event.Type), event)
}

func Subscribe(eventType EventType, fn interface{}) {
	eventBus.Subscribe(string(eventType), fn)
}

func Unsubscribe(eventType EventType, fn interface{}) {
	eventBus.Unsubscribe(string(eventType), fn)
}

// 初始化事件总线，进行事件发布者和订阅者的绑定
func InitEventBus() {
	Subscribe(UserSignedUp, HandleUserSignedUp)
	Subscribe(UserSignedUp, NewUserDeletedSubsrciber().HandleUserDeleted)
}
