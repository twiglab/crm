package poly

import "context"

type Input struct {
}

type Output struct {
}

type RuleDesc struct {
	Code string
	Type string
}

type Ruler interface {
	Exec(ctx context.Context, in *Input, out *Output) error
}

type RulerLoader interface {
	LoadRuler(ctx context.Context, rd RuleDesc) (Ruler, error)
}
