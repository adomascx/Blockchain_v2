package lib

import (
	"time"
)

type Block struct {
	// Header
	prev_block_hash   string
	timestamp         time.Time
	version           string
	difficulty_target int
	nonce             int
	merkle_root       string
	// Body
	transactions []Transaction
}

func NewBlock(transactions []Transaction, prev_block_hash, version string, difficulty_target, nonce int) (Block, error) {
	// TODO: implement constructor
	return Block{}, nil

}

// Calculate Merkle root of the block's transactions.
func (b Block) calculateMerkleRoot() string {
	// handle block w/o transactions
	if len(b.transactions) == 0 {
		return PHA256([]byte(""))
	}

	// initialize merkle tree with tx ids as leaves
	var level []string
	for _, tx := range b.transactions {
		level = append(level, tx.transactionId)
	}

	for len(level) > 1 {
		// if block contains odd number of txs, duplicate last tx id
		if len(level)%2 == 1 {
			lastIdx := len(level) - 1
			level = append(level, level[lastIdx])
		}

		// build higher merkle tree level
		var nextLevel []string
		// concat pairs of ids, then hash the result into new level
		for i := 0; i < len(level); i += 2 {
			combined := level[i] + level[i+1]
			nextLevel = append(nextLevel, PHA256([]byte(combined)))
		}
		level = nextLevel
	}

	return level[0]
}
