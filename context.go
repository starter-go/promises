package promises

import "context"

// Context 是为 Promise 提供的上下文
type Context interface {
	context.Context

	Worker() // todo ...

	Dispatcher() // todo ...
}

func NewContext() Context {
	return nil
}
