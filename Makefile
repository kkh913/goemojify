GOCMD        = go
GOBUILD      = $(GOCMD) build
GOCLEAN      = $(GOCMD) clean
GOTEST       = $(GOCMD) test
GOVET        = $(GOCMD) vet
GOGET        = $(GOCMD) get
GOX          = $(GOPATH)/bin/gox
GOGET        = $(GOCMD) get

GOX_ARGS     = -output="$(BUILD_DIR)/{{.Dir}}-{{.OS}}-{{.Arch}}" -osarch="linux/amd64 darwin/amd64"

BUILD_DIR    = build
BINARY_NAME  = goemojify

all: clean vet test build

build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v -ldflags "-X main.GitTag=`git describe --tags --abbrev=0`"

vet:
	${GOVET} ./...

test:
	${GOTEST} ./...

coverage:
	${GOTEST} -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/*

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

release:
	${GOGET} -u github.com/mitchellh/gox
	${GOX} -ldflags "${LD_FLAGS}" ${GOX_ARGS}
	shasum -a 512 build/* > build/sha512sums.txt

docker:
	docker build --rm --force-rm --no-cache -t kkh913/goemojify .

.PHONY: all vet test coverage clean build run release docker
