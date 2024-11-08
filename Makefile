.PHONY: tidy run

tidy:
	go mod tidy

run:
	@echo "Running cmd/main.go"
	go run cmd/main.go
