.SILENT:

build: 
	go build -o ./bin/app ./cmd/app/main.go

run: build
	./bin/app

swagger:
	swag init -q -g cmd/app/main.go

migration-generate:
	cd ./migrations && goose create $(or $(NAME),name) sql && cd ..

migration-status:
	goose status

migration-run:
	goose up

migration-revert:
	goose down

docker: 
	docker-compose -p moto-alerts up -d

fmt:
	go fmt ./...