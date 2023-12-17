package mypackage

import "fmt"

// ExportedFunction はエクスポートされた関数です。
func ExportedFunction() {
    fmt.Println("これはエクスポートされた関数です。")
}

// nonExportedFunction は非エクスポートされた関数です。
func nonExportedFunction() {
    
    fmt.Println("This is a non-exported function.")
}
