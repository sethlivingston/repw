COMMIT_HASH :=$(shell git rev-parse --short HEAD)
ifndef VERSION
	VERSION := $(shell git describe --tags --abbrev=0).${COMMIT_HASH}
endif

all: build

build:
	go build -ldflags "-X github.com/sethlivingston/repw/internal/config.Version=${VERSION}"