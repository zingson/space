package router

import (
	"io"
	"log"
	"net/http"
)

func ListenServer() {
	log.Println("listenserver......")

	// 设置路由规则
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "version 1")
	})

	// 使用默认的DefaultServeMux
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}