FROM golang:latest

WORKDIR /app
COPY ./app /app

RUN go mod init main \
  && go mod tidy \
  && go build

# コンテナに永続的な環境変数を定義します。コピーする対象 コンテナ内のコピー先で指定できます
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080


CMD ["go", "run", "main.go"]