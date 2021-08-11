package wf

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const TaskListName = "helloWorldGroup"

func init() {
	workflow.Register(SimpleWorkflow)
	activity.Register(SimpleActivity)
}

// SimpleActivity is a sample Cadence activity function that takes one parameter and
// returns a string containing the parameter value.
func SimpleActivity(ctx context.Context, value string) (string, error) {
	activity.GetLogger(ctx).Info("SimpleActivity called.", zap.String("Value", value))
	spew.Dump("Hello: ", value)
	return "Processed: " + value, nil
}

func SimpleWorkflow(ctx workflow.Context, value string) error {
	ao := workflow.ActivityOptions{
		// TaskList:               "sampleTaskList",
		ScheduleToCloseTimeout: time.Second * 10,
		ScheduleToStartTimeout: time.Second * 10,
		StartToCloseTimeout:    time.Second * 10,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	future := workflow.ExecuteActivity(ctx, SimpleActivity, value)
	var result string
	if err := future.Get(ctx, &result); err != nil {
		return err
	}
	workflow.GetLogger(ctx).Info("Done", zap.String("result", result))

	// step 2
	signalName := "SignalName"
	selector := workflow.NewSelector(ctx)
	var ageResult int
	for {
		signalChan := workflow.GetSignalChannel(ctx, signalName)
		selector.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, &ageResult)
			workflow.GetLogger(ctx).Info("Received age results from signal!", zap.String("signal", signalName), zap.Int("value", ageResult))
		})
		workflow.GetLogger(ctx).Info("Waiting for signal on channel.. " + signalName)
		// Wait for signal
		selector.Select(ctx)

		break
	}
	spew.Dump("after break")
	return nil
}
