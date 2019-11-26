EXECUTABLE = bot

all: build

build:
	go build -o bin/$(EXECUTABLE) cmd/main.go

clean:
	rm $(EXECUTABLE)
