package test

import (
	"testing"

	"github.com/adomascx/Blockchain_v2/lib"
)

func TestBlockInit(t *testing.T) {
	const (
		nonce      = 123456789
		version    = "0.2"
		difficulty = 3
		prevHash   = ""
	)

	testBlock := lib.NewBlock([]lib.Transaction{}, prevHash, version, difficulty, nonce)

	body := testBlock.GetBody()

}
