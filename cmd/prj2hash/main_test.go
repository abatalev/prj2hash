package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(".", getRoot(""))
	assertions.Equal("a", getRoot("a"))
}

func TestProcess(t *testing.T) {
	assertions := require.New(t)
	files, hash := process("", "../../examples/example/")
	assertions.Len(files, 2)
	assertions.Len(hash, 40)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", files[0].Hash)
	assertions.Equal("5441d4130251f67a2827b8a19122f6af0c4ceda7", files[1].Hash)
}
