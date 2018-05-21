package memo

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error) // 待缓存的函数类型

type result struct {
	// 函数的返回结果
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key) // 将函数结果缓存
		// 在两个临界区之间，多个goroutine可以并发的执行函数获取结果
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err // 不需要再次执行函数，直接返回缓存的结果
}
