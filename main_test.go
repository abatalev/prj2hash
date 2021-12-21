package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExcludeMask(t *testing.T) {

	data := []struct {
		cfg      config
		filename string
		result   bool
	}{
		{cfg: config{Excludes: []string{}}, filename: "a.txt", result: false},
		{cfg: config{Excludes: []string{"a\\.txt"}}, filename: "a.txt", result: true},
		{cfg: config{Excludes: []string{".settings/**/*"}}, filename: ".settings/a.txt", result: true},
		{cfg: config{Excludes: []string{".settings/**/*"}}, filename: ".settings/lib/a.txt", result: true},
		{cfg: config{Excludes: []string{"**/*.txt"}}, filename: "target/b.txt", result: true},
		{cfg: config{Excludes: []string{"target/*"}}, filename: "target/a.txt", result: true},
		{cfg: config{Excludes: []string{"target/**/*"}}, filename: "target/lib/a.txt", result: true},
		{cfg: config{Excludes: []string{"target/**/*.js"}}, filename: "target/lib/a.js", result: true},
		{cfg: config{Excludes: []string{"target/**/*.js"}}, filename: "target/lib/a.cs", result: false},
	}

	assertions := require.New(t)
	for _, variant := range data {
		assertions.Equal(variant.result, excludeMask(&variant.cfg, variant.filename), "error on processing %s", variant.filename)
	}
}

func TestLoadConfig(t *testing.T) {
	assertions := require.New(t)
	cfg := loadConfig("xxx.xxx")
	assertions.Equal(0, len(cfg.Excludes), "variant 1")
	cfg = loadConfig("./example/.prj2hash.yaml")
	fmt.Println(cfg)
	assertions.Equal(1, len(cfg.Excludes), "variant 2")
}
func TestCalcHashBytes(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", calcHashBytes([]byte("a")), "???")
}

func TestCalcHashFiles(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("90ce62edf2fe4940e041a68b13e7b5f9d02bbf51", calcHashFiles([]fileInfo{{fileName: "a", hash: "b"}}), "???")
}

func TestWalk(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(2, len(makeFileList(&config{Excludes: []string{}}, "example")), "length of example 1")
	assertions.Equal(1, len(makeFileList(&config{Excludes: []string{".prj2hash.yaml"}}, "example")), "length of example 2")
}

func TestSortFiles(t *testing.T) {
	assertions := require.New(t)
	files := sortFiles([]fileInfo{{"b", "1"}, {"a", "1"}})
	assertions.Len(files, 2)
	assertions.Equal("a", files[0].fileName)
	assertions.Equal("b", files[1].fileName)
}

func TestShortHash(t *testing.T) {
	assertions := require.New(t)
	assertions.Len(getShortHash("12345678901234567890", true), 8)
	assertions.Len(getShortHash("12345678901234567890", false), 20)
}

func TestRoot(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(".", getRoot(""))
	assertions.Equal("a", getRoot("a"))
}

func TestProcess(t *testing.T) {
	assertions := require.New(t)
	files, hash := process(".prj2hash.yaml", "./example/")
	assertions.Len(files, 1)
	assertions.Len(hash, 40)
}
