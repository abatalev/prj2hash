package files

import (
	"testing"

	"github.com/abatalev/prj2hash/internal/config"
	"github.com/stretchr/testify/require"
)

func TestWalk(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(3, len(MakeFileList(&config.Config{Excludes: []string{}},
		"../../examples/example")), "length of example 1")
	assertions.Equal(2, len(MakeFileList(&config.Config{Excludes: []string{".prj2hash.yaml"}},
		"../../examples/example")), "length of example 2")
}
