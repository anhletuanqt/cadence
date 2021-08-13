# How to start

1. Run docker

- `make docker_cp`

2. Register `test-domain` domain

- `make domain`

3. Run worker

- `make worker`

4. Run nats

- make sure nats is running on your local (port: 4222)
- `make nats`

5. Run server

- `make server`

6. Run web

- cd web && `yarn install` to install npm packages
- `make web`

# How it work

- Order worfkow has 4 activities: `Start Preparing`, `Notify Customer`, `Update LCD`, `Preparing Order`

- `Start Preparing`: wait for StartPreparingSignal signal
- `Notify Customer`, `Update LCD`: external task
- `Preparing Order`: wait for PreparingOrderSignal signal

- For simplicity, when you run the workflow, after 5s I will send `StartPreparingSignal` signal, and after the next 5s I will send `PreparingOrderSignal` signal

```go
	// Create workflow
	client, err := cadenceClient.CadenceClient.StartWorkflow(context.Background(), wo, order_workflow.OrderWorkFlow, orderID, customerID)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// Send StartPreparingSignal
	time.Sleep(5 * time.Second)
	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.StartPreparingSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// Send PreparingOrderSignal
	time.Sleep(5 * time.Second)
	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.PreparingOrderSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}
```

- localhost ClientUI: http://localhost:8088/domains/test-domain/workflows?range=last-30-days
