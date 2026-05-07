package lib

import (
	"errors"
	"math/rand/v2"
	"strconv"
)

func generateTransactions(size int) ([]Transaction, error) {

	return nil, nil
}

type user struct {
	name      string
	publicKey string
	balance   float64
}

func generateUsers(size int) ([]user, error) {
	if size < 1 {
		return nil, errors.New("Invalid user size. Should be more than 0")
	}

	generatedUsers := make([]user, size)
	for i := range size {
		generatedUsers = append(generatedUsers, user{
			name:      "bob" + strconv.Itoa(i+1),
			publicKey: PHA256([]byte(strconv.Itoa(i + 1))),
			balance:   (rand.Float64()*1e6 - 100) + 100,
		})
	}

	return generatedUsers, nil
}

func txToJson(transactions []Transaction) (string, error) {
	return "", nil
}
