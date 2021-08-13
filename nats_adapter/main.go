package main

import (
	cadenceAdapter "client/adapter"
	"client/config"
	"client/util"
	order_activity "client/worker/activity/order"
	order_workflow "client/worker/workflow/order"
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/cadence/client"
)

func main() {
	var appConfig config.AppConfig
	appConfig.Setup("../resources")
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)
	natsConn, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}

	// Create workflow
	natsConn.Subscribe(util.NATS_NEW_ORDER, func(msg *nats.Msg) {
		orderID := string(msg.Data)
		wo := client.StartWorkflowOptions{
			ID:                           orderID,
			TaskList:                     order_workflow.TaskListName,
			ExecutionStartToCloseTimeout: time.Hour * 24,
		}

		_, err := cadenceClient.CadenceClient.StartWorkflow(context.Background(), wo, order_workflow.OrderWorkFlow, orderID)
		if err != nil {
			msg.Respond([]byte(err.Error()))
		}

		msg.Respond([]byte{})
	})

	// Star preparing signal
	natsConn.Subscribe(util.NATS_START_PREPARING_ORDER, func(msg *nats.Msg) {
		orderID := string(msg.Data)
		err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.StartPreparingSignal, nil)
		if err != nil {
			msg.Respond([]byte(err.Error()))
		}

		msg.Respond([]byte{})
	})

	// Preparing order signal
	natsConn.Subscribe(util.NATS_PREPARING_ORDER, func(msg *nats.Msg) {
		orderID := string(msg.Data)
		err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.PreparingOrderSignal, nil)
		if err != nil {
			msg.Respond([]byte(err.Error()))
		}

		msg.Respond([]byte{})
	})

	select {}
}
