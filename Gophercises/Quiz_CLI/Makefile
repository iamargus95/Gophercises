.PHONY = all clean
all: build

build: quiz_cli.go
		@echo "Building Binary...."
		go build -o quiz_cli.o

run:
		./quiz_cli.o

clean: 
		@echo "Cleaning up..."
		rm quiz_cli.o
		go clean