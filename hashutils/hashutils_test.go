package hashutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcHashBytes(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", CalcHashBytes([]byte("a")), "???")
}

func TestCalcHashFiles(t *testing.T) {
	data := []FileInfo{{FileName: "a", Hash: "b"}}
	assertions := require.New(t)
	assertions.Equal("90ce62edf2fe4940e041a68b13e7b5f9d02bbf51", CalcHashFiles(data), "???")
}

func TestSortFiles(t *testing.T) {
	assertions := require.New(t)
	files := SortFiles([]FileInfo{{FileName: "b", Hash: "1"}, {FileName: "a", Hash: "2"}})
	assertions.Len(files, 2)
	assertions.Equal("a", files[0].FileName)
	assertions.Equal("2", files[0].Hash)
	assertions.Equal("b", files[1].FileName)
	assertions.Equal("1", files[1].Hash)
}
