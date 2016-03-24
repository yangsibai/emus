run: build
	./emus

build:
	go build -o emus

live:
	git push live

.PHONY: run build live
