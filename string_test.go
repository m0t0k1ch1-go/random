package randomutil_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/m0t0k1ch1-go/randomutil"
)

var (
	alphanumericRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func TestAlphanumericString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			s, err := randomutil.AlphanumericString(8)
			require.Nil(t, err)

			require.Len(t, s, 8)
			for _, r := range s {
				require.True(t, slices.Contains(alphanumericRunes, r))
			}
		}
	})
}
