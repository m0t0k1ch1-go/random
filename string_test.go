package random_test

import (
	"slices"
	"testing"

	"github.com/m0t0k1ch1-go/random"
	"github.com/m0t0k1ch1-go/random/internal/testutil"
)

var (
	alphanumericRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func TestAlphanumericString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			s, err := random.AlphanumericString(8)
			if err != nil {
				t.Fatal(err)
			}

			testutil.Equal(t, 8, len(s))
			for _, r := range s {
				testutil.Equal(t, true, slices.Contains(alphanumericRunes, r))
			}
		}
	})
}
