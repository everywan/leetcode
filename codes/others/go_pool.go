package others

import (
	"sync"
	"time"
)

type Pool struct {
	maxCap int32 // 最大容量
	chs    chan chan func()
	doneCh chan struct{} // 结束

	minCap      int32 // 最小容量
	currLen     int32
	maxIdelTime time.Duration // 最大空闲时间
}

func NewPool(maxCap int32) {
	pool := &Pool{
		chs: make(chan chan func(), maxCap),
	}
	pool.init()
}

func (pool *Pool) Run(fn func()) {
	select {
	case ch := <-pool.chs:
		ch <- fn
	case <-pool.doneCh:
		return
	default:
		// locker
		if pool.currLen < pool.maxCap {
			// 新增一个
			return
		}
		// 否则等待 chs 空闲
	}
}

func (pool *Pool) init() {
	eg := sync.WaitGroup{}
	eg.Add(int(pool.maxCap))
	for i := 0; i < int(pool.maxCap); i++ {
		// chanel 长度对数据不影响, 因为有执行任务时不会放到 chs 里
		ch := make(chan func())
		go func() {
			for {
				timer := time.NewTimer(pool.maxIdelTime)
				select {
				case fn := <-ch:
					fn()
					timer.Reset(pool.maxIdelTime)
					pool.chs <- ch
				case <-pool.doneCh:
					return
				case <-timer.C:
					if pool.currLen >= pool.minCap {
						pool.currLen--
						return
					}
					timer.Reset(pool.maxIdelTime)
				}
			}
		}()
		pool.chs <- ch
	}
	eg.Wait()
}

func (pool *Pool) Close() {
	close(pool.doneCh)
}
