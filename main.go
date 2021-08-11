package main

import (
	cadenceAdapter "client/adapter"
	"client/config"

	// wf "client/worker/workflow"

	"context"
	"fmt"
)

func main() {
	var appConfig config.AppConfig
	appConfig.Setup("./resources")
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)
	// wo := client.StartWorkflowOptions{
	// 	ID:                           "2",
	// 	TaskList:                     wf.TaskListName,
	// 	ExecutionStartToCloseTimeout: time.Hour * 24,
	// }
	// client, err := cadenceClient.CadenceClient.StartWorkflow(context.Background(), wo, wf.SimpleWorkflow, "Tuan Anh")
	// if err != nil {
	// 	fmt.Println("err: ", err)
	// }
	// spew.Dump("client:", client)

	err := cadenceClient.CadenceClient.SignalWorkflow(context.Background(), "2", "", "SignalName", 25)
	if err != nil {
		fmt.Println("err: ", err)
	}

}
