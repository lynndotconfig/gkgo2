package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，
// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

func ServeApp(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello, gopher")
	})
	addr := "0.0.0.0:8080"
	return serve(ctx, addr, mux)
}

func serveDebug(ctx context.Context) error {
	addr := "127.0.0.1:8001"
	return serve(ctx, addr, http.DefaultServeMux)
}

func serve(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:              addr,
		Handler:           handler,
	}
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
		log.Println(fmt.Sprintf("shutdown server: %s", addr))
	}()
	log.Printf("start server: %s", addr)
	return s.ListenAndServe()
}
func signalListener(ctx context.Context, ) error {
	notifyCh :=make(chan os.Signal, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)
	
	log.Println("awaiting signal")
	sign := <- notifyCh
	log.Printf("receive signal: %s", sign)
	return errors.New("process killed")
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return serveDebug(ctx)
	})
	g.Go(func() error {
		return ServeApp(ctx)
	})
	g.Go(func() error {
		return signalListener(ctx)
	})
	
	if err := g.Wait(); err != nil {
		log.Println("server error")
	}
}

// output below
// 2021/09/19 19:41:39 start server: 127.0.0.1:8001
// 2021/09/19 19:41:39 start server: 0.0.0.0:8080
// 2021/09/19 19:41:39 awaiting signal
// 2021/09/19 19:41:42 receive signal: interrupt
// 2021/09/19 19:41:42 shutdown server: 0.0.0.0:8080
// 2021/09/19 19:41:42 shutdown server: 127.0.0.1:8001
// 2021/09/19 19:41:42 server error