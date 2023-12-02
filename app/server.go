// go:generate go get -u package
package main

import (
	"html/template"
  "io"
  "net/http"
  "github.com/labstack/echo/v4" 
)

// Echoフレームワーク用のカスタムhtml/テンプレートレンダラ
type TemplateRenderer struct {
	templates *template.Template
}

// テンプレート文書をレンダリングする
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

  // データがマップの場合、グローバルメソッドを追加する
  if viewContext, isMap := data.(map[string]interface{}); isMap {
  	viewContext["reverse"] = c.Echo().Reverse
  }

  return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()
  renderer := &TemplateRenderer{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
  e.Renderer = renderer

  e.GET("/user", func(c echo.Context) error {
    return c.Render(http.StatusOK, "user.html", map[string]interface{}{
  	  "name": "Usaaer!",
    })
  }).Name = "foobar"

  e.GET("/post", func(c echo.Context) error {
    return c.Render(http.StatusOK, "post.html", map[string]interface{}{
  	  "name": "Post!",
    })
  }).Name = "foobar"

  e.GET("hello",Hello)

  e.Logger.Fatal(e.Start(":8080"))
}
// func Hello(c echo.Context) error {
// 	return c.Render(http.StatusOK, "hello", "Worldaaaa")
// }
