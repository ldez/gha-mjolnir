.PHONY: clean check test build

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

default: clean check test build

clean:
	rm -rf dist/ builds/ cover.out

build: clean
	@echo Version: $(VERSION)
	go build -v -ldflags '-X "main.version=${VERSION}" -X "main.commit=${SHA}" -X "main.date=${BUILD_DATE}"' -o mjolnir

test: clean
	go test -v -cover ./...

check:
	golangci-lint run

image:
	docker build -t ldez/gha-mjolnir:latest .
	docker tag ldez/gha-mjolnir:latest ldez/gha-mjolnir:${VERSION}

publish-image:
	docker push ldez/gha-mjolnir:latest
	docker push ldez/gha-mjolnir:${VERSION}
