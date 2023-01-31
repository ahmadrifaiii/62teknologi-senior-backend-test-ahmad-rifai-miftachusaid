REST_MAIN := "$(CURDIR)/cmd/rest"
BIN_REST := "$(CURDIR)/bin/rest-service"

init:
	@go mod init
	
tidy:
	@go mod tidy

fetch:
	@go mod download

build-rest:
	@go build -i -v -o $(BIN_REST) $(REST_MAIN)

build-rest-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_REST) $(REST_MAIN)

clean:
	@rm -rf $(CURDIR)/bin

deploy: clean fetch build-rest 

swagger-gen:
	@swag init -g cmd/rest/main.go --output pkg/swagger/docs

run-migrate:
	@go run utl/database/migrate/mysql/main.go