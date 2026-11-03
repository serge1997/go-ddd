package result

type Result[T any] struct {
	isSuccess bool
	value     T
	err       error
}

func Fail[T interface{}](err error) Result[T] {
	var z T
	return Result[T]{isSuccess: false, value: z, err: err}
}

func Ok[T interface{}](v T) Result[T] {
	return Result[T]{isSuccess: true, value: v, err: nil}
}

func (r Result[T]) Value() T {
	return r.value
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) IsFail() bool {
	return r.isSuccess == false
}
