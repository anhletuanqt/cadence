# How to start

1. Run cadence server with docker

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

- localhost ClientUI: http://localhost:8088/domains/test-domain/workflows?range=last-30-days
