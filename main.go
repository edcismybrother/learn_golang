package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// defer ants.Release()
	// runTimes := 10000
	// // use the common pool
	// // var wg sync.WaitGroup
	// for i := 0; i < runTimes; i++ {
	// 	// wg.Add(1)
	// 	ants.Submit(func() error {
	// 		handle("哈哈", i)
	// 		// wg.Done()
	// 		return nil
	// 	})
	// }
	// // wg.Wait()
	// fmt.Printf("running goroutines: %d\n", ants.Running())
	// fmt.Printf("finish all tasks.\n")

	for i := 0; i < 1000; i++ {
		// time.Sleep(2 * time.Millisecond)
		p.push(Msg{
			from: uint64(i),
			data: "haha",
		})
	}

	fmt.Println("参与的worker数量：", len(p.workers))
	time.Sleep(1 * time.Second)
}

type Pool struct {
	max     uint64
	workers []*Worker //存放工作者
	chl     chan Msg
	lock    sync.Mutex
}

type Msg struct {
	from uint64 //来自
	data string // 消息内容
}

type Worker struct {
	ID          uint64 //id
	pool        *Pool  //属于哪一个池
	h           chan Msg
	recycleTime time.Time //回收时间
	end         chan bool
}

var p = newPool(10)

var max uint64

func newPool(i uint64) *Pool {
	p := &Pool{
		max:     i,
		workers: make([]*Worker, 0),
		chl:     make(chan Msg, 1024),
	}
	go p.handle(i)
	return p
}

func (p *Pool) push(msg Msg) {
	p.chl <- msg
}

func (p *Pool) putWorker(w *Worker) {
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.lock.Unlock()
}

func (p *Pool) handle(bm uint64) {
	for {
		msg := <-p.chl
		var w *Worker
		if atomic.LoadUint64(&max) >= bm {
			for {
				l := len(p.workers) - 1
				if l < 0 {
					// p.lock.Unlock()
					continue
				} else {
					p.lock.Lock()
					ws := p.workers
					w = ws[l]
					ws[l] = nil
					p.workers = ws[:l]
					p.lock.Unlock()
					break
				}
			}
		} else if w == nil {
			next := atomic.AddUint64(&max, uint64(1))
			w = &Worker{
				ID:   next - 1,
				pool: p,
				end:  make(chan bool, 1),
				h:    make(chan Msg, 1),
			}
			go w.run()
		}
		w.h <- msg
	}
}

func (w *Worker) handle(msg Msg) {
	fmt.Printf("工人:%v执行:%v的%v\n", w.ID, msg.from, msg.data)
	// w.pool.lock.Lock()
	w.pool.putWorker(w)
	// w.pool.lock.Unlock()
}

func (w *Worker) run() {
	for {
		msg := <-w.h
		w.handle(msg)
	}
}

func handle(name string, i int) {
	fmt.Printf("工人:%v执行哈哈%v\n", name, i)
}
