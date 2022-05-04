server:
    go run cmd/api/server.go
build:
 	 go build -o bin/api cmd/api/server.go 
test:
		go test ./...
cover:
		go test ./... -cover
		go test tool cover -func profile.cov
		go tool cover -html profile.cov -o coverage.html