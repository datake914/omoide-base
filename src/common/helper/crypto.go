package helper

import (
	"crypto/sha256"
	"encoding/hex"

	hashids "github.com/speps/go-hashids"
)

func HashidsEncode(salt string, minlength int, numbers []int) (string, error) {
	hash := hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		Salt:      salt,
		MinLength: minlength,
	})
	return hash.Encode(numbers)
}

func HashidsEncodeBytes(salt string, minlength int, bytes []byte) (string, error) {
	i := []int{}
	for _, b := range bytes {
		i = append(i, int(b))
	}
	return HashidsEncode(salt, minlength, i)
}

func Sha256(bytes []byte) string {
	checksum := sha256.Sum256(bytes)
	return hex.EncodeToString(checksum[:])
}
