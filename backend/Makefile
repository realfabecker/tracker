.PHONY: build-docs
build-docs:
	swag init --parseDependency --parseInternal -g cmd/http/main.go -o ./internal/handlers/http/docs/

.PHONY: build-http
build-http:
	GOOS=linux GOARCH=amd64 go build -o ./out/http ./cmd/http/main.go

.PHONY: build-lambda
build-lambda:
	GOOS=linux GOARCH=arm64 go build -o ./out/bootstrap ./cmd/lambda/main.go