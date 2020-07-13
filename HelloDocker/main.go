package main

import (
	"fmt"
	"log"
	"net/http"
)

// 8080ポートでサーバーアプリケーションとして動作
// クライアントからリクエストを受けた場合は標準出力に表示
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request")
		fmt.Fprint(w, "Hello Docker!!")		
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}