package main

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SimpleWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, SimpleActivity, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func SimpleActivity(ctx context.Context, name string) (string, error) {
	return "Hello " + name + "!", nil
}
