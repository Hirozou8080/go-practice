package main

import (
	"net/http"
	"main/mypackage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

  // ルートを設定
  e.GET("/", hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

  // サーバーをポート番号1323で起動
  e.Logger.Fatal(e.Start(":8080"))

  // エクスポートされた関数を呼び出す
  mypackage.ExportedFunction()

  // 非エクスポートされた関数は他のファイルからは直接呼び出せない
  // mypackage.nonExportedFunction()  // コンパイルエラーになります
}

// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
