package order_activity

import (
	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/workflow"
)

const StartPreparingSignal = "StartPreparingSignal"

func StartPreparing(ctx workflow.Context, orderID string) error {
	signalName := StartPreparingSignal
	selector := workflow.NewSelector(ctx)
	for {
		signalChan := workflow.GetSignalChannel(ctx, signalName)
		selector.AddReceive(signalChan, func(c workflow.Channel, more bool) {
			c.Receive(ctx, nil)
		})
		// Wait for signal
		selector.Select(ctx)

		spew.Dump("PreparingOrder: order ", orderID)

		return nil
	}
}
