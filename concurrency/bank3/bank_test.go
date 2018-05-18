package bank_test

import (
	"testing"
	"sync"
	"../bank3"
)

func TestBank(t *testing.T) {
	var n sync.WaitGroup
	// 1000个存款goroutine并发执行
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			bank.Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()
	if got, want := bank.Balance(), (1+1000)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
