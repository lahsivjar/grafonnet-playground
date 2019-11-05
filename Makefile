all: clean test build

APP=grafonnet-playground
ALL_PACKAGES=$(shell go list ./...)
SOURCE_DIRS=$(shell go list ./... | cut -d "/" -f4 | uniq)

setup:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/fzipp/gocyclo
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/client9/misspell
	go get -u github.com/alecthomas/gometalinter

tidy:
	GO111MODULE=on go mod tidy -v

clean: tidy
	rm -rf ./dist
	rm -rf ./out

check-quality: lint fmt cyclo vet

go-build:
	GO111MODULE=on go build -o "./out/${APP}"

go-release-build:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o "./out/${APP}"

npm-build-dev:
	npm run dev

npm-build:
	npm run build

build-dev: npm-build-dev go-build

build: npm-build go-build

build-release: npm-build go-release-build

test:
	GO111MODULE=on go test -count 1 -cover -v ./...

fmt:
	gofmt -l -s -w $(SOURCE_DIRS)

imports:
	go get -u golang.org/x/tools/cmd/goimports
	goimports -l -w -v $(SOURCE_DIRS)

cyclo:
	go get -u github.com/fzipp/gocyclo
	gocyclo -over 7 $(SOURCE_DIRS)

vet:
	GO111MODULE=on go vet ./...

lint:
	go get -u golang.org/x/lint/golint
	golint -set_exit_status ./...

race:
	GO111MODULE=on go test -race -short ./...

msan: # Memory sanitizer (will only work in linux/amd64)
	GO111MODULE=on go test -msan -short ./...

coverage:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	# Below displays global test coverage
	GO111MODULE=on go tool cover -func=coverage.out

coverage-html:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	# Below displays global test coverage
	GO111MODULE=on go tool cover -func=coverage.out
	mkdir public
	GO111MODULE=on go tool cover -html=./coverage.out -o public/coverage.html

copy-config:
	cp application.yml.sample application.yml

