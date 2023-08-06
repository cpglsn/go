## build_go_front-end: build front-end go code
build_go_front-end:
	@echo "Building front-end go code"
	cd front-end && env CGO_ENABLED=0 go build -o front-end ./cmd/web
	@echo "front-end go code compiled"


## build_container_front-end: build front-end container
build_container_front-end:
	@echo "Building front-end go container"
	docker compose build front-end
	@echo "front-end container built"


## build_go_broker-service: build broker-service go code
build_go_broker-service:
	@echo "Building broker-service go code"
	cd broker-service && env GOOS=linux CGO_ENABLED=0 go build -o broker-service ./cmd/api
	@echo "broker-service go code compiled"


## build_container_broker-service: build broker-service container
build_container_broker-service:
	@echo "Building broker-service go container"
	docker compose build broker-service
	@echo "broker-service container built"


## build_go_authentication-service: build authentication-service go code
build_go_authentication-service:
	@echo "Building authentication-service go code"
	cd authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o authentication-service ./cmd/api
	@echo "authentication-service go code compiled"


## build_container_authentication-service: build authentication-service container
build_container_authentication-service:
	@echo "Building authentication-service go container"
	docker compose build authentication-service
	@echo "authentication-service container built"


## build_go_all: build all go code
build_go_all: build_go_front-end build_go_broker-service build_go_authentication-service
	@echo "All go code compiled"


## build_container_all: build all containers
build_container_all: build_container_front-end build_container_broker-service build_container_authentication-service
	@echo "All containers built"


## build: build all
build: build_go_all build_container_all
	@echo "All built"


## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker compose up -d
	sleep 10s
	PGPASSWORD=secret psql -h localhost -U postgres -p 5432 -d users < users.sql
	@echo "Docker images started!"


## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"
