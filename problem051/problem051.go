package problem051

import (
	"crypto/rand"
	"math/big"
)

func whatAPerfectRandom(k int) (int, error) {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(k)))
	if err != nil {
		return 0, err
	}
	return int(randomInt.Int64() + 1), err
}

func swap(values *[]int, indexA int, indexB int) {
	(*values)[indexA], (*values)[indexB] = (*values)[indexB], (*values)[indexA]
}

func Shuffle(values *[]int) error {
	for i := range *values {
		j, err := whatAPerfectRandom(len(*values) - i)
		if err != nil {
			return err
		}
		swap(values, i, j-1+i)
	}
	return nil
}
