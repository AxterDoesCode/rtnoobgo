all:
	make build exec
build:
	go build -o bin/ ./cmd/rt
exec:
	./bin/rt

