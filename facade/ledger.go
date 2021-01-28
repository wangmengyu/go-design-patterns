package facade

import "fmt"

/**
账本
*/
type Ledger struct {
}

func (l *Ledger) MakeEntry(accountId string, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountId, txnType, amount)
	return
}
