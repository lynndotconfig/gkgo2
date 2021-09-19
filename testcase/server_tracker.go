package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Tracker struct {
	ch chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch:   make(chan string, 10),
	}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <- ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run() {
	for data := range t.ch {
		time.Sleep(1 * time.Second)
		fmt.Println(data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown(ctx context.Context) error {
	// 写的一方决定什么时候，channel可以关闭
	// 关闭写channle
	close(t.ch)
	
	// 等待runner发会关闭信号
	select {
	case <-t.stop:
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
}

func main () {
	tr := NewTracker()
	go tr.Run()
	ctx := context.Background()
	_ = tr.Event(ctx, "test1")
	_ = tr.Event(ctx, "test2")
	_ = tr.Event(ctx, "test3")
	_ = tr.Event(ctx, "test4")
	time.Sleep(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5 * time.Second))
	defer cancel()
	tr.Shutdown(ctx)
}