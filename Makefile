build:
	go get -v github.com/mtojek/go-keylogger/cmd/keylogger

clean-code:
	goimports -w .
	golint ./... && test -z "$$?"

test: build
	go get -t ./...
	keylogger version
	keylogger devices
	# keylogger record

pre-push: clean-code build test
