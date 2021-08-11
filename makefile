docker_cp:
	cd docker && docker compose up

domain:
	docker run --network=host --rm ubercadence/cli:master --do test-domain domain register -rd 1

worker: 
	cd worker && go run worker.go

workflow:
	go run main.go

.PHONY: docker_cp domain worker workflow