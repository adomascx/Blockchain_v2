package test

import (
	"bufio"
	"fmt"
	"os"
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

func TestHashDictionary(t *testing.T) {
	// read dictionary text file into string array
	dictPath := "words_2.txt"
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
		t.Errorf("Hash() = %v duplicates detected:\n%v", nDuplicates, dupes)
	}

}
