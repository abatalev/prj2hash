package rules

import (
	"path/filepath"
	"strings"

	"github.com/abatalev/prj2hash/internal/config"
	"github.com/bmatcuk/doublestar/v4"
)

type Rule struct {
	Allow bool
	Mask  string
}

func Convert(cfg *config.Config) []string {
	if len(cfg.Rules) > 0 {
		return cfg.Rules
	}

	xRules := []string{"allow **/*"}
	for _, str := range cfg.Excludes {
		xRules = append(xRules, "deny "+str)
	}
	return xRules
}

func ConvertRulesToStruct(rulesStrings []string) []Rule {
	rules := []Rule{}
	for _, str := range rulesStrings {
		if strings.HasPrefix(str, "allow ") {
			rules = append(rules, Rule{Allow: true, Mask: strings.TrimPrefix(str, "allow ")})
		} else {
			rules = append(rules, Rule{Allow: false, Mask: strings.TrimPrefix(str, "deny ")})
		}
	}
	return rules
}

func CheckFileByRules(rules []Rule, path string) bool {
	path0 := filepath.ToSlash(path)
	mm := false
	result := false
	for _, rule := range rules {
		matched, _ := doublestar.Match(rule.Mask, path0)
		if matched {
			mm = true
			result = rule.Allow
		}
	}
	if mm {
		return !result
	}

	return true
}
