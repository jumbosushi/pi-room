.PHONY: bin
bin: main.go sensors/*.go
	go build -o bin/pi-room main.go
