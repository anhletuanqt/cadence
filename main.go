package main

import (
	cadenceAdapter "client/adapter"
	"client/config"
	order_activity "client/worker/activity/order"
	order_workflow "client/worker/workflow/order"
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/client"
)

func main() {
	var appConfig config.AppConfig
	appConfig.Setup("./resources")
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)
	customerID := gofakeit.UUID()
	orderID := gofakeit.UUID()
	wo := client.StartWorkflowOptions{
		ID:                           orderID,
		TaskList:                     order_workflow.TaskListName,
		ExecutionStartToCloseTimeout: time.Hour * 24,
	}
	client, err := cadenceClient.CadenceClient.StartWorkflow(context.Background(), wo, order_workflow.OrderWorkFlow, orderID, customerID)
	if err != nil {
		fmt.Println("err: ", err)
	}
	spew.Dump("client:", client)

	time.Sleep(5 * time.Second)
	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.StartPreparingSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}
	time.Sleep(5 * time.Second)

	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.PreparingOrderSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}

}
