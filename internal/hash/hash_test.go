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
