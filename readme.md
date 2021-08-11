# How to start

1. Run docker

- `make docker_cp`

2. Register `test-domain` domain

- `make domain`

3. Run worker

- `make worker`

4. Run sample workflow

- `make workflow`

# How it work

- Order worfkow has 4 activities: `Start Preparing`, `Notify Customer`, `Update LCD`, `Preparing Order`

- `Start Preparing`: wait for signal StartPreparingSignal
- `Notify Customer`, `Update LCD`: external task
- `Preparing Order`: wait for signal PreparingOrderSignal

- For simplicity, when you run the workflow, after 5s I will send `StartPreparingSignal` signal, and after the next 5s I will send `PreparingOrderSignal`

```go
	// Create workflow
	client, err := cadenceClient.CadenceClient.StartWorkflow(context.Background(), wo, order_workflow.OrderWorkFlow, orderID, customerID)
	if err != nil {
		fmt.Println("err: ", err)
	}
	spew.Dump("client:", client)

	// Send StartPreparingSignal
	time.Sleep(5 * time.Second)
	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.StartPreparingSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}
	time.Sleep(5 * time.Second)

	// Send PreparingOrderSignal
	err = cadenceClient.CadenceClient.SignalWorkflow(context.Background(), orderID, "", order_activity.PreparingOrderSignal, nil)
	if err != nil {
		fmt.Println("err: ", err)
	}
```

- localhost ClientUI: http://localhost:8088/domains/test-domain/workflows?range=last-30-days
