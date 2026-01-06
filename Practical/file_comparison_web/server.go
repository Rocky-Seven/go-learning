package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	
	fmt.Println("サーバー起動: http://localhost:8080")
	fmt.Println("file-diff.html にアクセスしてください")
	http.ListenAndServe(":8080", nil)
}
