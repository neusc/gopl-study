package bank

import "sync"

var (
	mu      sync.Mutex // 监控变量
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock() // 即使发生panic，defer语句依然可以执行
	return balance
}
