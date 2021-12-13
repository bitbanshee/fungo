package primitives

type Observable interface {
	Subscribe(id string, ob ObserverFunc)
	Unsubscribe(id string)
}

type ObserverFunc func(interface{})
