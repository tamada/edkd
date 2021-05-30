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
  
  test: setup
	  $(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

  build: setup
	  $(GO) build -o $(NAME) cmd/edkd/*.go
    
  dist: build
    @$(call _createDist,darwin,amd64,)
    @$(call _createDist,darwin,386,)
    @$(call _createDist,windows,amd64,.exe)
    @$(call _createDist,windows,386,.exe)
    @$(call _createDist,linux,amd64,)
    @$(call _createDist,linux,386,)
    
  clean:
  	$(GO) clean
