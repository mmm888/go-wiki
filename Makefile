GO     := $(shell which go)

BIN := go-wiki
APP := github.com/mmm888/go-wiki


.PHONY:	build vendor clean run


build: vendor
	go build -o $(BIN) .


vendor: dep-install
ifneq ($(VENDOR_MD5),$(GOPKG_MD5))
	dep ensure
ifneq ($(shell type md5 2> /dev/null),)
	md5 -q Gopkg.lock >| vendor/lock.md5
else ifneq ($(shell type md5sum 2> /dev/null),)
	md5sum Gopkg.lock | sed -E 's/ .*//g' >| vendor/lock.md5
else
	@echo vendor/lock.md5 was not created 1>&2
endif
else
	@echo vendor/ is already up-to-date
endif


clean:
	rm -f $(BIN)
	rm -rf vendor
	rm -rf $(GOPATH)/pkg/linux_amd64/$(APP)/*
	rm -rf $(GOPATH)/pkg/darwin_amd64/$(APP)/*
	dep remove


dep-install:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif


dep-ensure: dep-install
	dep ensure


run: vendor
	go run *.go
