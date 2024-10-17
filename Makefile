.PHONY: all
all:
	go mod vendor -v

.PHONY: run
run:
	go run main.go