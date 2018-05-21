package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error) // 待缓存的函数类型

type result struct {
	// 函数的返回结果
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // res准备好之后关闭该channel
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 第一次请求某个Url时
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		// 对某个url的重复请求
		memo.mu.Unlock()
		// 重复请求的goroutine必须等待ready之后才能读取条目的结果
		// 在channel关闭之前此处一直阻塞
		<-e.ready
	}

	return e.res.value, e.res.err // 不需要再次执行函数，直接返回缓存的结果
}
