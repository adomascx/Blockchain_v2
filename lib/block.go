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
	panic("TODO: implement constructor")
}

// Calculate Merkle root of the block's transactions.
func (b Block) calculateMerkleRoot() string {
	if len(b.transactions) == 0 {
		return PHA256([]byte(""))
	}

	var level []string
	for _, tx := range b.transactions {
		level = append(level, tx.transactionId)
	}

	for len(level) > 1 {
		if len(level)%2 == 1 {
			level = append(level, level[len(level)-1])
		}

		var nextLevel []string
		for i := 0; i < len(level); i += 2 {
			combined := level[i] + level[i+1]
			nextLevel = append(nextLevel, PHA256([]byte(combined)))
		}
		level = nextLevel
	}

	return level[0]
}

// getter for block's header.
func (b Block) GetHeader() map[string]string {

}
