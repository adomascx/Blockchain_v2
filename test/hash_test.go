package test

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/adomascx/Blockchain_v2/lib"
)

func ReadDictFile(path string) ([]string, error) {
	var words []string

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not do os.Open(%v): %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}

func TestHashCollisions(t *testing.T) {
	// Print hash with the error or not
	const verbose = false
	// read dictionary text file into string array
	const dictPath = "words_alpha.txt"
	words, err := ReadDictFile(dictPath)
	if err != nil {
		t.Errorf("could not do ReadDictFile(%v): %v", dictPath, err)
	}

	// a reverse hash map (hash -> word)
	// this makes duplicate retrieval easier
	hashMap := make(map[string]string)

	// map of duplicates (hash -> duplicate words array)
	dupes := make(map[string][]string)

	for _, word := range words {
		// hash the word
		hash := lib.PHA256([]byte(word))

		// if it's in the hashmap already, it's a duplicate
		if hashMap[hash] != "" {
			if len(dupes[hash]) == 0 {
				dupes[hash] = append(dupes[hash], hashMap[hash])
			}

			dupes[hash] = append(dupes[hash], word)
		}

		hashMap[hash] = word
	}

	nDuplicates := len(dupes)
	if nDuplicates > 0 {
		var errorStr strings.Builder

		errorStr.WriteString(fmt.Sprintf("Hash() - %v hash collisions detected:\n", nDuplicates))

		for hash, inputs := range dupes {
			if verbose {
				errorStr.WriteString(fmt.Sprintf("Hash: %v, inputs: %v\n", hash, inputs))
			} else {
				errorStr.WriteString(fmt.Sprintf("Inputs: %v\n", inputs))
			}
		}

		t.Errorf("%s", errorStr.String())

	}

}

func TestHashLen(t *testing.T) {
	testStrings := []string{"a", "b", "abd", "", strings.Repeat("X", 256)}
	const expectedLen = 64

	for _, str := range testStrings {
		length := len(lib.PHA256([]byte(str)))
		if length != expectedLen {
			t.Errorf("Hash(%q) - Length of hash is not %v: len = %v\n", str, expectedLen, length)
		}
	}
}

func TestHashDeterminism(t *testing.T) {
	const testString = "asdf"
	hash1 := lib.PHA256([]byte(testString))
	hash2 := lib.PHA256([]byte(testString))

	if hash1 != hash2 {
		t.Errorf("Hash(%q) - Same input gives different outputs (non-deterministic):\nHash 1: %q\nHash 2: %q\n", testString, hash1, hash2)
	}
}

func TestHashAvalanche(t *testing.T) {
	const bitSize = 256             // no. of bits produced by hash
	const requiredDifference = 0.35 // bare minimum percentage
	const goodDifference = 0.47     // not exactly 50 due to statistical variance

	testString1 := strings.Repeat("a", 12) + "b"
	testString2 := strings.Repeat("a", 12) + "c"
	hash1 := lib.PHA256([]byte(testString1))
	hash2 := lib.PHA256([]byte(testString2))

	testHash1 := new(big.Int)
	testHash1.SetString(hash1, 16)

	testHash2 := new(big.Int)
	testHash2.SetString(hash2, 16)

	xor := new(big.Int)
	xor.Xor(testHash1, testHash2)

	diff := 0
	for i := 0; i < bitSize; i++ {
		if xor.Bit(i) == 1 {
			diff++
		}
	}

	expectedDiff := bitSize * requiredDifference
	if diff < int(expectedDiff) {
		t.Errorf("Hash() - Avalanche effect non-existent: hamming distance of %v\n", diff)
	}

	expectedDiff = bitSize * goodDifference
	if diff < int(expectedDiff) {
		t.Errorf("Hash() - Avalanche effect weak: hamming distance of %v\n", diff)
	}
}
