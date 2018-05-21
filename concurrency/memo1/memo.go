package memo

type Memo struct {
	f     Func
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
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key) // 将函数结果缓存
		memo.cache[key] = res
	}
	return res.value, res.err // 不需要再次执行函数，直接返回缓存的结果
}
