package main

import (
	"client/config"

	cadenceAdapter "client/adapter"
	order_workflow "client/worker/workflow/order"

	"github.com/davecgh/go-spew/spew"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

func startWorkers(h *cadenceAdapter.CadenceAdapter, taskList string) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}

	cadenceWorker := worker.New(h.ServiceClient, h.Config.Domain, taskList, workerOptions)
	err := cadenceWorker.Start()
	if err != nil {
		h.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}

func main() {
	spew.Dump("Starting Worker..")

	var appConfig config.AppConfig
	appConfig.Setup("../resources")

	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)

	startWorkers(&cadenceClient, order_workflow.TaskListName)
	select {}
}
