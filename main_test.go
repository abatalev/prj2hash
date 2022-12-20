package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	assertions := require.New(t)
	cfg := loadConfig("xxx.xxx")
	assertions.Equal(0, len(cfg.Excludes), "variant 1")
	cfg = loadConfig("./examples/example/.prj2hash.yaml")
	assertions.Equal(1, len(cfg.Excludes), "variant 2")
}

func TestCalcHashBytes(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", calcHashBytes([]byte("a")), "???")
}

func TestCalcHashFiles(t *testing.T) {
	data := []fileInfo{{fileName: "a", hash: "b"}}
	assertions := require.New(t)
	assertions.Equal("90ce62edf2fe4940e041a68b13e7b5f9d02bbf51", calcHashFiles(data), "???")
}

func TestWalk(t *testing.T) {
	assertions := require.New(t)
	assertions.Equal(3, len(makeFileList(&config{Excludes: []string{}}, "examples/example")), "length of example 1")
	assertions.Equal(2, len(makeFileList(&config{Excludes: []string{".prj2hash.yaml"}}, "examples/example")), "length of example 2")
}

func TestSortFiles(t *testing.T) {
	assertions := require.New(t)
	files := sortFiles([]fileInfo{{"b", "1"}, {"a", "2"}})
	assertions.Len(files, 2)
	assertions.Equal("a", files[0].fileName)
	assertions.Equal("2", files[0].hash)
	assertions.Equal("b", files[1].fileName)
	assertions.Equal("1", files[1].hash)
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
	files, hash := process("", "./examples/example/")
	assertions.Len(files, 2)
	assertions.Len(hash, 40)
	assertions.Equal("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", files[0].hash)
	assertions.Equal("5441d4130251f67a2827b8a19122f6af0c4ceda7", files[1].hash)
}

func TestConvertExcludesToRules(t *testing.T) {
	data := []struct {
		cfg   config
		rules []string
	}{
		{cfg: config{Excludes: []string{}}, rules: []string{"allow **/*"}},
		{cfg: config{Excludes: []string{"a\\.txt"}}, rules: []string{"allow **/*", "dany a\\.txt"}},
	}

	assertions := require.New(t)
	for idx, variant := range data {
		xRules := convert(&variant.cfg)
		assertions.Equal(len(variant.rules), len(xRules),
			"error on processing %d", idx)
	}
}

func TestRules(t *testing.T) {
	data := []struct {
		cfg      config
		filename string
		result   bool
	}{
		{cfg: config{Excludes: []string{}}, filename: "a.txt", result: false},
		{cfg: config{Excludes: []string{"a\\.txt"}}, filename: "a.txt", result: true},
		{cfg: config{Rules: []string{"allow **/*", "deny a\\.txt"}}, filename: "a.txt", result: true},
	}

	assertions := require.New(t)
	for _, variant := range data {
		assertions.Equal(variant.result,
			checkFileByRules(convertRulesToStruct(convert(&variant.cfg)), variant.filename),
			"error on processing %s", variant.filename)
	}
}
