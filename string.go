package randomutil

import (
	"crypto/rand"
	"math/big"

	"github.com/pkg/errors"
)

var (
	alphanumericRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alphanumericMax   = big.NewInt(int64(len(alphanumericRunes)))
)

// AlphanumericString returns a random alphanumeric string of the given length.
func AlphanumericString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("length must not be negative")
	}

	rs := make([]rune, length)
	{
		for i := 0; i < length; i++ {
			n, err := rand.Int(rand.Reader, alphanumericMax)
			if err != nil {
				return "", errors.Wrap(err, "failed to read random int")
			}

			rs[i] = alphanumericRunes[n.Int64()]
		}
	}

	return string(rs), nil
}
