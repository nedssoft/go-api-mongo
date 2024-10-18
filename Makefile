build:
	@go build -o bin/go-api-mongo cmd/main.go


test:
	@go test -v ./...



run: build
	@./bin/go-api-mongo



