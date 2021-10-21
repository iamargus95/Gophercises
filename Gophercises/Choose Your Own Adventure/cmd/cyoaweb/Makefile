.PHONY = all clean
all: build

build: main.go
		@echo "Building Binary...."
		go build

run:
		make build
		./cyoaweb

clean: 
		@echo "Cleaning up..."
		rm cyoaweb
		go clean