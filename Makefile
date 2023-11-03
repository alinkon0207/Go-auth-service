all:
	go mod tidy
	swag init
	go run main.go