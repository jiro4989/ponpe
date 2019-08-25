APPNAME := $(shell basename `pwd`)
VERSION := $(shell grep Version version.go | grep -Eo "version.*" | grep -Eo '[\.0-9]+$$')
LDFLAGS := -ldflags="-s -w \
	-extldflags \"-static\""
XBUILD_TARGETS := \
	-os="windows linux darwin" \
	-arch="386 amd64" 
DIST_DIR := dist
README := README.*
EXTERNAL_TOOLS := \
	github.com/mitchellh/gox

.PHONY: help
help: ## ドキュメントのヘルプを表示する。
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## ビルド
	go build $(LDFLAGS) -o bin/$(APPNAME) .

.PHONY: install
install: ## インストール
	go install $(LDFLAGS)

.PHONY: xbuild
xbuild:  bootstrap ## クロスコンパイル
	gox $(LDFLAGS) $(XBUILD_TARGETS) --output "$(DIST_DIR)/{{.Dir}}$(VERSION)_{{.OS}}_{{.Arch}}/{{.Dir}}"

.PHONY: archive
archive: xbuild ## クロスコンパイルしたバイナリとREADMEを圧縮する
	find $(DIST_DIR)/ -mindepth 1 -maxdepth 1 -a -type d \
		| while read -r d; \
		do \
			cp $(README) $$d/ ; \
			cp LICENSE $$d/ ; \
		done
	cd $(DIST_DIR) && \
		find . -maxdepth 1 -mindepth 1 -a -type d  \
		| while read -r d; \
		do \
			../tools/archive.sh $$d; \
		done

.PHONY: test
test: ## テストコードを実行する
	go test -cover ./...

.PHONY: clean
clean: ## バイナリ、配布物ディレクトリを削除する
	-rm -rf bin
	-rm -rf $(DIST_DIR)

.PHONY: bootstrap
bootstrap: ## 外部ツールをインストールする
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		GO111MODULE=off go get $$t ; \
	done
