build: clean-code install test

clean-code:
	go get golang.org/x/tools/cmd/goimports && goimports -w .
	go get golang.org/x/lint/golint && golint -set_exit_status ./...

install:
	go get -v github.com/mtojek/go-keylogger/cmd/keylogger

test: install
	go get -t ./...
	keylogger || test -n "$$?"
	keylogger version --help
	keylogger devices --help
	keylogger record --help

vagrant-deploy:
	GOOS=linux GOARCH=amd64 go build -v github.com/mtojek/go-keylogger/cmd/keylogger
	vagrant scp keylogger :/home/vagrant
	rm keylogger
