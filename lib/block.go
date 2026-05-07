package lib

import (
	"strconv"
	"strings"
	"time"
)

type Block struct {
	header       header        // Header
	transactions []Transaction // Body
}

type header struct {
	prevBlockHash    string
	timestamp        time.Time
	version          string
	difficultyTarget int
	nonce            int
	merkleRoot       string
}

func NewBlock(transactions []Transaction, prevBlockHash, version string, difficultyTarget, nonce int) *Block {
	b := new(Block)

	b.transactions = transactions
	b.header.version = version
	b.header.difficultyTarget = difficultyTarget
	b.header.nonce = nonce
	b.header.timestamp = time.Now()
	b.header.merkleRoot = b.calculateMerkleRoot()

	if prevBlockHash == "" {
		b.header.prevBlockHash = PHA256([]byte(""))
	}
	b.header.prevBlockHash = prevBlockHash

	return b
}

// Calculate Merkle root of the Block's transactions.
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

// getter for Block's header.
func (b Block) GetHeader() header {
	return b.header
}

// getter for Block's transactions.
func (b Block) GetBody() []Transaction {
	return b.transactions
}

// Calculate the Block's hash based on its header.
func (b *Block) calculateHash() string {
	headerParts := []string{
		b.header.prevBlockHash,
		b.header.timestamp.String(),
		b.header.version,
		b.header.merkleRoot,
		strconv.Itoa(b.header.nonce),
		strconv.Itoa(b.header.difficultyTarget),
	}
	return PHA256([]byte(strings.Join(headerParts, "|")))
}
