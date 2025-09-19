package files

import (
	"testing"

	"abatalev.com/prj2hash/internal/config"
	"github.com/stretchr/testify/require"
)

func TestSortFiles(t *testing.T) {
	assertions := require.New(t)
	files := SortFiles([]FileInfo{{"b", "1"}, {"a", "2"}})
	assertions.Len(files, 2)
	assertions.Equal("a", files[0].FileName)
	assertions.Equal("2", files[0].Hash)
	assertions.Equal("b", files[1].FileName)
	assertions.Equal("1", files[1].Hash)
}

func TestWalk(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(3, len(MakeFileList(&config.Config{Excludes: []string{}},
		"../../examples/example")), "length of example 1")
	assertions.Equal(2, len(MakeFileList(&config.Config{Excludes: []string{".prj2hash.yaml"}},
		"../../examples/example")), "length of example 2")
}

func TestCalcHashFiles(t *testing.T) {
	data := []FileInfo{{FileName: "a", Hash: "b"}}
	assertions := require.New(t)
	assertions.Equal("90ce62edf2fe4940e041a68b13e7b5f9d02bbf51", CalcHashFiles(data), "???")
}
