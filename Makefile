.PHONY: default bindata fmt

TAGS = release
PORTAL_DATA_FILES := $(shell find portal | sed 's/  /\\ /g')
ASSETS_DATA_FILES := $(shell find assets | sed 's/  /\\ /g')
GENERATED := pkg/assets/bindata.go pkg/portal/bindata.go

LDFLAGS += -X "github.com/alimy/gin-music/version.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/alimy/gin-music/version.BuildGitHash=$(shell git rev-parse HEAD)"

default: all

all: fmt
	go build -ldflags '$(LDFLAGS)' -tags '$(TAGS)' github.com/alimy/gin-music/cmd/gin-music

fmt:
	go fmt github.com/alimy/gin-music/...

bindata: $(GENERATED)
	gofmt -s -w pkg/assets pkg/portal

pkg/assets/bindata.go: $(ASSETS_DATA_FILES)
	rm -rf pkg/assets/
	go-bindata -nomemcopy -pkg=assets -tags=$(TAGS) \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ assets/...

pkg/portal/bindata.go: $(PORTAL_DATA_FILES)
	rm -rf pkg/portal
	go-bindata -nomemcopy -pkg=assets  -tags=$(TAGS) \
         -debug=$(if $(findstring debug,$(TAGS)),true,false) \
         -o=$@ portal/...

clean:
	go clean -r github.com/alimy/gin-music/...
	rm -f gin-music

distclean: clean
	rm -rf pkg/assets/ pkg/portal/