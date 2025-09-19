package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	assertions := require.New(t)
	cfg := LoadConfig("xxx.xxx")
	assertions.Equal(0, len(cfg.Excludes), "variant 1")
	cfg = LoadConfig("../../examples/example/.prj2hash.yaml")
	assertions.Equal(1, len(cfg.Excludes), "variant 2")
}
