package memo

type Memo struct {
	requests chan request
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

type request struct {
	key      string
	response chan<- result
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

// 将map变量限制在一个单独的监控goroutine中
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 第一次请求计算函数返回结果
		}
		go e.deliver(req.response) // 之后的重复调用直接从cache返回结果
	}
}

// 第一次调用函数的计算过程
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready) // 向其它goroutine广播该条目内容已经设置完成的信息
}

//
func (e *entry) deliver(response chan<- result) {
	<-e.ready // 等待第一次函数调用完成之后的广播消息
	response <- e.res
}
