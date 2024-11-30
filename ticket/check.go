package ticket

import "context"

type Input struct {
}

type Output struct {
}

type CheckRuler interface {
	func(ctx context.Context, in *Input, out *Output) error
}
