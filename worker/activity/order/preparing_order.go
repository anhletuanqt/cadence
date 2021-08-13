package order_activity

import (
	"client/server/db"

	"go.uber.org/cadence/workflow"
)

// func init() {
// 	activity.Register(PreparingOrder)
// }

// func PreparingOrder(ctx context.Context, orderID string, customerID string) (string, error) {
// 	// activity.GetLogger(ctx).Info("SimpleActivity called.", zap.String("Value", value))
// 	spew.Dump("PreparingOrder: order ", orderID, "customer:", customerID)
// 	return orderID, nil
// }

const PreparingOrderSignal = "PreparingOrderSignal"

func PreparingOrder(ctx workflow.Context, orderID string) error {
	signalName := PreparingOrderSignal
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
		order.Activities = append(order.Activities, "preparing order")
		// set order
		db.SetOrderByID(order)

		return nil
	}
}
