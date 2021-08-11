package order_workflow

import (
	order_activity "client/worker/activity/order"
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
)

const TaskListName = "OrderWorkFlow"

func init() {
	workflow.Register(OrderWorkFlow)
}

func OrderWorkFlow(ctx workflow.Context, orderID string, customerID string) error {
	ao := workflow.ActivityOptions{
		// TaskList:               "sampleTaskList",
		ScheduleToCloseTimeout: time.Second * 10,
		ScheduleToStartTimeout: time.Second * 10,
		StartToCloseTimeout:    time.Second * 10,
		HeartbeatTimeout:       time.Second * 10,
		WaitForCancellation:    false,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Start Preparing
	if err := order_activity.StartPreparing(ctx, orderID); err != nil {
		return err
	}
	fmt.Println("==========StartPreparing=========")

	// Notify customer
	future := workflow.ExecuteActivity(ctx, order_activity.NotifyCustomer, orderID, customerID, "Start Preparing")
	var result1 string
	if err := future.Get(ctx, &result1); err != nil {
		return err
	}
	fmt.Println("==========NotifyCustomer=========")

	// Update lcd
	future = workflow.ExecuteActivity(ctx, order_activity.UpdateLCD, orderID, "StartPreparing")
	var result2 string
	if err := future.Get(ctx, &result2); err != nil {
		return err
	}
	fmt.Println("==========UpdateLCD=========")

	// Preparing Order
	if err := order_activity.PreparingOrder(ctx, orderID); err != nil {
		return err
	}
	fmt.Println("==========PreparingOrder=========")

	fmt.Println("Done")
	return nil
}
