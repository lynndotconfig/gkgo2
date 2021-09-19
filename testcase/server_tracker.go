package main

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type Tracker struct {
	wg sync.WaitGroup
}

func (t *Tracker) Event(data string) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		time.Sleep(time.Millisecond)
		log.Printf(data)
	}()
}

func (t *Tracker) Shutdown(ctx context.Context) error {
	ch := make(chan struct{})
	go func() {
		t.wg.Wait()
		close(ch)
	}()
	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("timeout")
	}
}

type App struct {
	t Tracker
}

func (a *App) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	a.t.Event("created")
}

func main () {
	// 启动一个服务
	var a App
	
	// 关闭服务
	
	// 等待所有事件的协程退出
	a.t.Shutdown()
}