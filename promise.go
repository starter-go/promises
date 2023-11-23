package promises

// ThenHandler 是 Promise.Then 方法的回调函数
type ThenHandler[T any] func(T) Promise[T]

// CatchHandler 是 Promise.Catch 方法的回调函数
type CatchHandler[T any] func(error) Promise[T]

// FinallyHandler 是 Promise.Finally 方法的回调函数
type FinallyHandler[T any] func() Promise[T]

// Promise 是 Promise 机制的主要接口
type Promise[T any] interface {
	Then(fn ThenHandler[T]) Promise[T]
	Catch(fn CatchHandler[T]) Promise[T]
	Finally(fn FinallyHandler[T]) Promise[T]
}

// Reject 返回一个已拒绝（rejected）的 Promise 对象，拒绝原因为给定的参数。
func Reject[T any](c Context, err error) Promise[T] {

	return nil // todo ...
}

// Resolve 方法将给定的值转换为一个 Promise。
func Resolve[T any](c Context, value T) Promise[T] {

	return nil // todo ...
}

func Race[T any]() Promise[T] {

	return nil // todo ...
}

func All[T any]() Promise[T] {

	return nil // todo ...
}

func Any[T any]() Promise[T] {

	return nil // todo ...
}

// Execute 执行指定任务
func Execute[T any](c Context, task func() (T, error)) Promise[T] {
	return nil // todo ...
}
