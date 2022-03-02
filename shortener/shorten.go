package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// hashThis receives string and return a slice of bytes
func hashThis(inputStr string) []byte {
	shortenAlgorithm := sha256.New()
	shortenAlgorithm.Write([]byte(inputStr))
	return shortenAlgorithm.Sum(nil)
}

// Encoder encodes byte and return url string
func Encoder(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := hashThis(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := Encoder([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
