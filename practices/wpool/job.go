package wpool

import (
	"context"
)

type JobID string
type jobType string
type jobMetadata map[string]interface{}

type ExecutionFn func(ctx context.Context, args interface{}) (interface{}, error)

type JobDescriptor struct {
	ID       JobID
	JType    jobType
	Metadata jobMetadata
}

type Result struct {
	Value      interface{}
	Err        error
	Descriptor JobDescriptor
}

type Job struct {
	Descriptor JobDescriptor
	ExecFn     ExecutionFn // It is defined which func will be executed in a job.
	Args       interface{}
}

func (j Job) execute(ctx context.Context) Result {
	value, err := j.ExecFn(ctx, j.Args)
	if err != nil {
		return Result{
			Err:        err,
			Descriptor: j.Descriptor,
		}
	}

	return Result{
		Value:      value,
		Descriptor: j.Descriptor,
	}
}
