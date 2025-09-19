package rules

import (
	"testing"

	"abatalev.com/prj2hash/internal/config"
	"github.com/stretchr/testify/require"
)

func TestConvertExcludesToRules(t *testing.T) {
	data := []struct {
		cfg   config.Config
		rules []string
	}{
		{cfg: config.Config{Excludes: []string{}}, rules: []string{"allow **/*"}},
		{cfg: config.Config{Excludes: []string{"a\\.txt"}}, rules: []string{"allow **/*", "dany a\\.txt"}},
	}

	assertions := require.New(t)
	for idx, variant := range data {
		v := variant
		xRules := Convert(&v.cfg)
		assertions.Equal(len(variant.rules), len(xRules),
			"error on processing %d", idx)
	}
}

func TestRules(t *testing.T) {
	data := []struct {
		cfg      config.Config
		filename string
		result   bool
	}{
		{cfg: config.Config{Excludes: []string{}}, filename: "a.txt", result: false},
		{cfg: config.Config{Excludes: []string{"a\\.txt"}}, filename: "a.txt", result: true},
		{cfg: config.Config{Rules: []string{"allow **/*", "deny a\\.txt"}}, filename: "a.txt", result: true},
		{cfg: config.Config{Rules: []string{"allow **/*", "deny target/**/*"}}, filename: "target\\a.txt", result: false},
		{cfg: config.Config{Rules: []string{"deny **/*", "allow target/**/*"}}, filename: "target\\a.txt", result: true},
	}

	assertions := require.New(t)
	for _, variant := range data {
		v := variant
		assertions.Equal(variant.result,
			CheckFileByRules(ConvertRulesToStruct(Convert(&v.cfg)), variant.filename),
			"error on processing %s", variant.filename)
	}
}
