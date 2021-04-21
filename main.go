package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ葉１つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// ルート
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// Webサーバーの開始とエラーハンドリング
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
