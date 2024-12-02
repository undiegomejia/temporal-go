package dsl

import "go.temporal.io/sdk/workflow"

type (
	Workflow struct {
		Variables map[string]string
		Root      Statement
	}
	Statement struct {
		Activity *ActivityInvocation
		Sequence *Sequence
		Parallel *Parallel
	}
	Sequence struct {
		Elements []*Statement
	}
	Parallel struct {
		Branches []*Statement
	}
	ActivityInvocation struct {
		Name      string
		Arguments []string
		Result    string
	}
	executable interface {
		execute(ctx workflow.Context, bindings map[string]string) error
	}
)
