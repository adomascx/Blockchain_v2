package lib

// Represents a transaction with I/O and ID (UTXO model).
type Transaction struct {
	// TODO: implement correctly
	// inputs map[string]string
	// outputs map[string]string
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
