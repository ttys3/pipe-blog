# Makefile for golang app build
# author: 荒野无灯
# date: 2019-10-22

APP_NAME := pipe
GOPROXY := https://goproxy.cn,direct
# go-sqlite3 requires cgo to work. This is a stub
CGO_ENABLED := 1
GOX_OSARCH := linux/amd64 linux/arm64 darwin/amd64 windows/amd64
PKG_URL := github.com/b3log/pipe/model
APP_VERSION := $(shell cat build.version)
AUTO_VERSIONING := -X $(PKG_URL).Version=$(APP_VERSION)

define colorecho
      @TERM=xterm-256color tput setaf 6
      @echo $1
      @TERM=xterm-256color tput sgr0
endef

prep:
	test -x "$$(command -v rice)" || GO111MODULE=on go install github.com/GeertJohan/go.rice/rice
	test -x "$$(command -v gox)" || GO111MODULE=on go install github.com/mitchellh/gox

all: ui-admin ui-theme debug

release: ui-admin ui-theme linux

crossbuild: prep
	$(call colorecho,"remove release dir ...")
	rm -rf release
	mkdir -p ./release/console
	mkdir -p ./release/theme
	$(call colorecho,"begin build release binary ...")
	cd ./release/ && GO111MODULE=on CGO_ENABLED=$(CGO_ENABLED) gox -osarch="$(GOX_OSARCH)" -output "$(APP_NAME)_{{.OS}}_{{.Arch}}" -ldflags "-s -w $(AUTO_VERSIONING)" ../
	$(call colorecho,"begin compress binary with upx ...")
	upx ./release/$(APP_NAME)_*
	cp -r ./console/dist ./release/console/
	rsync -aqP --exclude=node_modules --exclude=scss --exclude=gulpfile.js  --exclude=package.json ./theme/ ./release/theme
	cp pipe.json ./release/pipe.json
	$(call colorecho,"begin generate md5sum ...")
	md5sum ./release/$(APP_NAME)_* > ./release/$(APP_NAME).md5sum.txt
	ls -lhp ./release
	$(call colorecho,"done.")

debug:
	GOPROXY=$(GOPROXY) go build -i -v .

win:
	GO111MODULE=on CGO_ENABLED=1 gox -mod="vendor" -osarch="windows/amd64" -output "filebrowser_{{.OS}}_{{.Arch}}" -ldflags "-s -w $(AUTO_VERSIONING)"

linux:
	GOPROXY=$(GOPROXY) go build -i -ldflags "-s -w $(AUTO_VERSIONING)" .

ui-admin:
	cd console && yarn install && yarn run build

ui-theme:
	cd theme && yarn install && yarn run build

dev:
	cd console && yarn install && yarn run dev

clean:
	rm -f ./$(APP_NAME)
	rm -rf ./console/dist
	rm -rf ./theme/dist
	rm -rf ./release