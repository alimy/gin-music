.PHONY: default build bindata fmt clean distclean

TAGS = release
PORTAL_DATA_FILES := $(shell find portal | sed 's/  /\\ /g')
ASSETS_DATA_FILES := $(shell find assets | sed 's/  /\\ /g')
GENERATED := pkg/assets/bindata.go pkg/portal/bindata.go

LDFLAGS += -X "github.com/alimy/gin-music/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/gin-music/version.GitHash=$(shell git rev-parse HEAD)"

default: build

build: fmt bindata
	go build -ldflags '$(LDFLAGS)' -tags '$(TAGS)'

fmt:
	go fmt ./...

bindata: $(GENERATED)

pkg/assets/bindata.go: $(ASSETS_DATA_FILES)
	rm -rf pkg/assets/
	go-bindata -nomemcopy -pkg=assets \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ assets/...
	gofmt -s -w pkg/assets


pkg/portal/bindata.go: $(PORTAL_DATA_FILES)
	rm -rf pkg/portal
	go-bindata -nomemcopy -pkg=assets  -tags=$(TAGS) \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ portal/...
	gofmt -s -w pkg/portal

clean:
	go clean -r ./...
	rm -f gin-music

distclean: clean
	rm -rf pkg/assets/ pkg/portal/