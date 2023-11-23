package unit

import (
	"testing"

	"github.com/starter-go/promises"
)

func TestPromise(t *testing.T) {

	ctx := promises.NewContext()

	promises.Execute[int](ctx, func() (int, error) {
		return 777, nil

	}).Then(func(i int) promises.Promise[int] {
		return promises.Resolve[int](ctx, 666)

	}).Catch(func(err error) promises.Promise[int] {
		return promises.Reject[int](ctx, err)

	}).Finally(func() promises.Promise[int] {
		return nil
	})

}
