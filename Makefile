GO     := $(shell which go)

BIN := go-wiki
APP := github.com/mmm888/go-wiki


.PHONY:	build vendor clean run test


build: vendor assets
	go build -o $(BIN) .


# TODO: md5 とって変更ないときは dep ensure を実行しない
vendor: dep-install
	dep ensure


assets: packr


clean:
	rm -f $(BIN)
	rm -rf vendor
	rm -rf $(GOPATH)/pkg/linux_amd64/$(APP)/*
	rm -rf $(GOPATH)/pkg/darwin_amd64/$(APP)/*
	dep remove
	packr clean


dep-install:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif


dep-ensure: dep-install
	dep ensure


packr: packr-install
	packr


packr-install:
ifeq ($(shell type packr 2> /dev/null),)
	go get -u github.com/gobuffalo/packr/...
endif


run: vendor assets
	go run *.go


test: 
	go test $(APP)/... -cover
