
# 変数定義
TEST = 'テスト'

# コマンド make [target]　-> targetがないと1番上のみが実行される。
.PHONY: all
ALL = 'buildされました'
all: 
	$(shell go mod init main) 
	$(shell go mod tidy) 
	$(shell go build -o bin/main)
	@echo $(ALL)


# buildされたファイルの削除
.PHONY: clean
CLEAN = 'クリーンされました'
clean: 
	$(shell rm bin/main)
	@echo $(CLEAN)

.PHONY: run
run:
	$(shell bin/main)
