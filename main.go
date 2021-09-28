package main

import (
	"log"
	"net/http"
	"time"

	"github.com/xueqiya/zheng_server/router"
)

func main() {
	server := &http.Server{
		Addr:           "0.0.0.0:8888",
		Handler:        router.Setup(),
		ReadTimeout:    time.Duration(60 * int(time.Second)), // 转换成时间数据结构
		WriteTimeout:   time.Duration(60 * int(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
