GO := go
NAME := edkd
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup: update_version

update_version:
	@for i in README.md docs/content/_index.md ; do \
		sed -e 's!Version-[0-9.]*-blue!Version-${VERSION}-blue!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' cmd/$(NAME)/main.go > a
	@mv a cmd/$(NAME)/main.go
	@sed 's/ARG version=.*/ARG version=${VERSION}/g' Dockerfile > b
	@mv b Dockerfile
	@echo "Replace version to \"${VERSION}\""
	
define __create_dist()
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) main.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
	tar cvfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef
	
dist: all
	@$(call __create_dist,darwin,amd64,)
	@$(call __create_dist,darwin,arm64,)
	@$(call __create_dist,windows,amd64,.exe)
	@$(call __create_dist,linux,amd64,)
  
test: setup
  	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)
	
build: setup
  	$(GO) build -o $(NAME) cmd/edkd/*.go

clean:
	$(GO) clean
