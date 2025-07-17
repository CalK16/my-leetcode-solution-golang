all: test
test:
	go test ./...

clean:
	rm __debug_*
