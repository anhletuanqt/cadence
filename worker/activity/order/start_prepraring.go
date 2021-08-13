package order_activity

import (
	"client/server/db"

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

		// get order
		order := db.GetOrderByID(orderID)
		order.Activities = append(order.Activities, "start preparing")
		// set order
		db.SetOrderByID(order)

		return nil
	}
}
