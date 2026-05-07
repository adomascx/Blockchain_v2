package lib

// Represents a transaction with I/O and ID (UTXO model).
type Transaction struct {
	inputs        []string
	outputs       []map[string]any
	transactionId string

	// input

}
type UTXO struct {
	value float64
	id    string
}

type Wallet struct {
	address string
	utxos   []UTXO
}

func newTransaction(owner string, amount int) Transaction {
	return Transaction{}
}

// TODO: make explicit param type
func (t Transaction) fromDict(payload any) Transaction {
	return Transaction{}
}
