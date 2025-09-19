package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortHash(t *testing.T) {
	assertions := require.New(t)
	assertions.Len(GetShortHash("12345678901234567890", true), 8)
	assertions.Len(GetShortHash("12345678901234567890", false), 20)
}

func TestCalcHashBytes(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", CalcHashBytes([]byte("a")), "???")
}
