package bank

var deposits = make(chan int) // 用于存款操作的channel
var balances = make(chan int) // 用于读取总金额的channel

func Deposit(amount int) { deposits <- amount } // 存款操作
func Balance() int       { return <-balances }  // 获取总金额

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits: // 存款
			balance += amount
		case balances <- balance: // 将总金额传入channel
		}
	}
}

func init() {
	go teller() // 启动监控goroutine
}
