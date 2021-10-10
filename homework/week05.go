package main

import (
	"sync"
	"time"
)

// 1. 参考 Hystrix 实现一个滑动窗口计数器。

type MetricCollector struct {
	mutex *sync.RWMutex
	numRequests int64
	errors int64
	
	success int64
	failures int64
	rejects int64
	shortCircuits int64
	timeouts int64
	
	runDuration int64
}

type Number struct {
	Buckets map[int64]*numberBucket	// 以 unixtimestamp 为 key
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

func (r *Number) getCurrentBucket() *numberBucket {
	// 先得到当前的 timestamp
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	// 判断是否存在，不存在则创建
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	
	return bucket
}

func (r *Number) removeOldBuckets() {
	now := time.Now().Unix() - 10
	
	for timestamp := range r.Buckets {
		if timestamp <= now {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}
	
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	// 先得到 bucket（当前 timestamp）
	b := r.getCurrentBucket()  //代码在上面，先以timestamp检查，不存在则新建
	b.Value += i
	// 删除掉旧的
	r.removeOldBuckets()
}

func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)
	
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	
	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-10 {
			// 确定 timestamp 在区间内
			sum += bucket.Value
		}
	}
	
	return sum
}