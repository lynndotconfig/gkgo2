package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func ServeApp(stop <- chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello, gopher")
	})
	addr := "0.0.0.0:8080"
	return serve(addr, mux, stop)
}

func serveDebug(stop <- chan struct{}) error {
	addr := "127.0.0.1:8001"
	return serve(addr, http.DefaultServeMux, stop)
}

func serve(addr string, handler http.Handler, stop <- chan struct{}) error {
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
	}
	go func() {
		// 接收一个channel的数据，结果不关心
		<-stop // 等待接收stop信号
		s.Shutdown(context.Background())
	}()
	log.Printf("listen: %s", addr)
	return s.ListenAndServe()
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- ServeApp(stop)
	}()
	
	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}