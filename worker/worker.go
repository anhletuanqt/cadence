package main

import (
	"client/config"
	"fmt"

	cadenceAdapter "client/adapter"
	wf "client/worker/workflow"

	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

// import (
// 	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
// 	"go.uber.org/cadence/worker"

// 	"github.com/uber-go/tally"
// 	"go.uber.org/yarpc"
// 	"go.uber.org/yarpc/transport/tchannel"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// )

// var HostPort = "127.0.0.1:7933"
// var Domain = "test-domain"
// var TaskListName = "simpleworker"
// var ClientName = "simpleworker"
// var CadenceService = "cadence-frontend"

// func main() {
// 	startWorker(buildLogger(), buildCadenceClient())
// }

// func buildLogger() *zap.Logger {
// 	config := zap.NewDevelopmentConfig()
// 	config.Level.SetLevel(zapcore.InfoLevel)

// 	var err error
// 	logger, err := config.Build()
// 	if err != nil {
// 		panic("Failed to setup logger")
// 	}

// 	return logger
// }

// func buildCadenceClient() workflowserviceclient.Interface {
// 	ch, err := tchannel.NewChannelTransport(tchannel.ServiceName(ClientName))
// 	if err != nil {
// 		panic("Failed to setup tchannel")
// 	}
// 	dispatcher := yarpc.NewDispatcher(yarpc.Config{
// 		Name: ClientName,
// 		Outbounds: yarpc.Outbounds{
// 			CadenceService: {Unary: ch.NewSingleOutbound(HostPort)},
// 		},
// 	})
// 	if err := dispatcher.Start(); err != nil {
// 		panic("Failed to start dispatcher")
// 	}

// 	return workflowserviceclient.New(dispatcher.ClientConfig(CadenceService))
// }

// func startWorker(logger *zap.Logger, service workflowserviceclient.Interface) {
// 	// TaskListName identifies set of client workflows, activities, and workers.
// 	// It could be your group or client or application name.
// 	workerOptions := worker.Options{
// 		Logger:       logger,
// 		MetricsScope: tally.NewTestScope(TaskListName, map[string]string{}),
// 	}

// 	worker := worker.New(
// 		service,
// 		Domain,
// 		TaskListName,
// 		workerOptions)
// 	err := worker.Start()
// 	if err != nil {
// 		panic("Failed to start worker")
// 	}

// 	logger.Info("Started Worker.", zap.String("worker", TaskListName))
// 	select {}
// }

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
	fmt.Println("Starting Worker..")
	var appConfig config.AppConfig
	appConfig.Setup("../resources")
	var cadenceClient cadenceAdapter.CadenceAdapter
	cadenceClient.Setup(&appConfig.Cadence)

	startWorkers(&cadenceClient, wf.TaskListName)
	// The workers are supposed to be long running process that should not exit.
	select {}
}
