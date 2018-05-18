package bank

var (
	sema    = make(chan struct{}, 1) // 二元信号量，控制在同一时刻最多只有一个goroutine访问一个共享变量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // 获取token
	balance += amount
	<-sema // 释放token
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
