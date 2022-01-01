FORCE:

http: FORCE
	@echo " >> building http binaries..."
	@go build -o cmd/http/http cmd/http/main.go
	@echo " >> http built."
	@echo "executing http..."
	@./cmd/http/http
	@echo "http is running."