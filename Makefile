GOPATH_DETECT=$(shell go env GOPATH)
GO_SKIP_NONIMPORTANT=go list ./... | grep -v vendor/ | grep -v mocks/ | xargs -L1

default: lint mods test sonar build
in-docker: lint mods test build

.PHONY: lint
lint:
	go install golang.org/x/lint/golint@latest
	$(GO_SKIP_NONIMPORTANT) $(GOPATH_DETECT)/bin/golint -set_exit_status

.PHONY: mods
mods:
	go mod tidy
	go mod download

.PHONY: test
test:
	$(GO_SKIP_NONIMPORTANT) go test -v -cover -coverprofile=bin/coverage.out

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build --ldflags="-s -w" -o bin/gopress-windows.exe ./cmd/gopress
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags="-s -w" -o bin/gopress-linux ./cmd/gopress

.PHONY: build-docker
build-docker:
	docker build -t mateuszmierzwinski/gopress:latest .

.PHONY: sonar
sonar:
	sonar-scanner
