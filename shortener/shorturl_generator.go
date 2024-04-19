package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func getHash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func generateNumber(input []byte) uint64 {
	return new(big.Int).SetBytes(input).Uint64()
}

func encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

/*
################## Algorithm ###################

 1. Hashing  initialUrl + userId url with sha256. Here userId
    is added to prevent providing similar shortened urls to
    separate users in case they want to shorten exact same link.

 2. Derive a big integer number from the hash bytes generated
    during the hashing.

 3. Finally apply base58  on the derived big integer value and
    pick the first 8 characters
*/
func GenerateShortLink(url string, userId string) string {
	urlHashBytes := getHash(url + userId)
	generatedNumber := generateNumber(urlHashBytes)
	finalString := encode([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
