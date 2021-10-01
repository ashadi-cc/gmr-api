server: 
	go run cmd/api/main.go
apidoc:
	swag init -g api/router.go