package util

import (
	"crypto/rand"
	"math/big"
)

var alphaNumeric = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func GenerateUniqueID(length int) string {
	randomString := ""
	maximum := big.NewInt(int64(len(alphaNumeric)))

	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, maximum)
		randomString += alphaNumeric[n.Int64()]
	}

	return randomString
}
