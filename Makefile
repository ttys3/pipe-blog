# Makefile for golang app build
# author: 荒野无灯
# date: 2019-10-22

APP_NAME := nanoblog
GOPROXY := https://goproxy.cn,direct
# go-sqlite3 requires cgo to work. This is a stub
CGO_ENABLED := 1
#GOX_OSARCH := linux/amd64 linux/arm64 darwin/amd64 windows/amd64
PKG_URL := github.com/b3log/pipe/model
APP_VERSION := $(shell cat build.version)
AUTO_VERSIONING := -X $(PKG_URL).Version=$(APP_VERSION)

define colorecho
      @TERM=xterm-256color tput setaf 6
      @echo $1
      @TERM=xterm-256color tput sgr0
endef

all: ui-admin ui-theme debug

release: ui-admin ui-theme linux

prep:
	test -x "$$(command -v rice)" || GO111MODULE=on go install github.com/GeertJohan/go.rice/rice

docker: prep ui-admin ui-theme
	$(call colorecho,"remove release dir ...")
	rm -rf release
	mkdir -p ./release/console
	mkdir -p ./release/theme
	mkdir -p ./release/i18n
	mkdir -p ./release/storage
	$(call colorecho,"begin build release binary ...")
	cd ./release/ && GO111MODULE=on CGO_ENABLED=$(CGO_ENABLED) go build -o "$(APP_NAME)" -ldflags "-s -w $(AUTO_VERSIONING)" ../
	$(call colorecho,"begin compress binary with upx ...")
	upx ./release/$(APP_NAME)
	cp -r ./console/dist ./release/console/
	rsync -aqP --exclude=node_modules --exclude='*.scss' --exclude=gulpfile.js --exclude=.idea --exclude=package.json ./theme/ ./release/theme
	rm -rf ./release/theme/js
	rm -rf ./release/theme/scss/animate.css
	cp -r ./i18n ./release/
	cp ./storage/.gitignore ./release/storage/
	cp pipe.json ./release/pipe.json
	sed -i 's/"RuntimeMode": "dev",/"RuntimeMode": "release",/' ./release/pipe.json
	sed -i 's/"LogLevel": "debug",/"LogLevel": "info",/' ./release/pipe.json
	$(call colorecho,"begin generate md5sum ...")
	md5sum ./release/$(APP_NAME) > ./release/$(APP_NAME).md5sum.txt
	ls -lhp ./release
	$(call colorecho,"done.")

debug:
	GOPROXY=$(GOPROXY) go build -o "$(APP_NAME)" -i -v .

win:
	GO111MODULE=on CGO_ENABLED=1 go build -o "$(APP_NAME)" -ldflags "-s -w $(AUTO_VERSIONING)"

linux:
	GOPROXY=$(GOPROXY) go build -o "$(APP_NAME)" -i -ldflags "-s -w $(AUTO_VERSIONING)" .

ui-admin:
	cd console && yarn install --verbose && yarn run build

ui-theme:
	cd theme && yarn install --verbose && yarn add aplayer && yarn run build

dev:
	cd console && yarn install --verbose && yarn run dev

clean:
	rm -f ./$(APP_NAME)
	rm -rf ./console/dist
	rm -rf ./release
	find ./theme -type f -name "*.min.*" ! -path './node_modules/*' | xargs rm -f