build:
	go get -v github.com/mtojek/go-keylogger/cmd/keylogger

clean-code:
	goimports -w .
	golint ./... && test -z "$$?"

test: build
	go get -t ./...
	keylogger version
	keylogger version --help
	keylogger devices
	keylogger devices --help
	# keylogger record
	keylogger record --help

pre-push: clean-code build test
