.PHONY: all
.PHONY: build
.PHONY: clean

# 変数定義
TEST = 'テスト'

# コマンド make [target]　-> targetがないと1番上のみが実行される。
all:
	$(shell go mod init main)
	$(shell go mod tidy)
	$(shell go build -o bin/main)


# buildされたファイルの削除
# clean:
#   $(shell rm bin/main)

build:
	$(shell bin/main)
