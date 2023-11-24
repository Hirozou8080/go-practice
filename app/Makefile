
# 変数定義
TEST = 'テスト'

# コマンド make [target]　-> targetがないと1番上のみが実行される。
.PHONY: all
all: 
	$(shell go mod init main) 
	$(shell go mod tidy) 
	$(shell go build -o bin/main)


# buildされたファイルの削除
.PHONY: clean
clean: 
	$(shell rm bin/main)

.PHONY: build
build:
	$(shell bin/main)
