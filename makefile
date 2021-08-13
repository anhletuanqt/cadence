docker_cp:
	cd docker && docker compose up

domain:
	docker run --network=host --rm ubercadence/cli:master --do test-domain domain register -rd 1

wk: 
	cd worker && go run worker.go

nats: 
	cd nats_adapter && go run main.go

server:
	cd server && go run *.go

web:
	cd web && yarn start

.PHONY: docker_cp domain worker server nats