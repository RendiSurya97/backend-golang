FORCE:

http: FORCE
	@echo " >> building http binaries..."
	@go build -o cmd/http/http cmd/http/main.go
	@echo " >> http built."
	@echo "executing http..."
	@./cmd/http/http
	@echo "http is running."

grpc: FORCE
	@echo " >> building grpc binaries..."
	@go build -o cmd/grpc/grpc cmd/grpc/main.go
	@echo " >> grpc built."
	@echo "executing grpc..."
	@./cmd/grpc/grpc
	@echo "grpc is running."

.DEFAULT_GOAL := http