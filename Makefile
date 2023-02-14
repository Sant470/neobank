build:
	@go build -o ./bin/neobank
run: build 
	@./bin/neobank
test:
	@go test -v ./...