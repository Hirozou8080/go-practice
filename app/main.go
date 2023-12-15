package main

import "main/mypackage"

func main() {
    // エクスポートされた関数を呼び出す
    mypackage.ExportedFunction()

    // 非エクスポートされた関数は他のファイルからは直接呼び出せない
    // mypackage.nonExportedFunction()  // コンパイルエラーになります
}
