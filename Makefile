run: build
	./emus

build:
	go build -o emus

.PHONY: run
