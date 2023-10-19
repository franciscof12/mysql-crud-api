build:
	@go build -o bin/personal_crud_api

run: build
	@./bin/personal_crud_api

test:
	@go test -v ./...

up:
	@docker-compose up
	
