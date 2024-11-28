package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func WorkflowOne(ctx workflow.Context, face FaceType) (FaceType, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 100,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	err := workflow.ExecuteActivity(ctx, AddHair, face).Get(ctx, &face.Hair)
	if err != nil {
		return face, err
	}

	err = workflow.ExecuteActivity(ctx, AddVoice, face).Get(ctx, &face.Voice)
	if err != nil {
		return face, err
	}

	return face, err
}
