build:
	go build -o school-auth ./cmd/main.go
sqlc:
	sqlc generate
swag:
	swag init -g /pkg/app/handler/server.go
up:
	docker-compose -f ./deployment/docker-compose.yaml up
up-build:
	docker-compose -f ./deployment/docker-compose.yaml up --build