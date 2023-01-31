package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"gopkg.in/yaml.v3"
)

type config struct {
	Excludes []string `yaml:"excludes"` // DEPRECATED! Remove in next version
	Rules    []string `yaml:"rules"`
}

type rule struct {
	Allow bool
	Mask  string
}

func convert(cfg *config) []string {
	if len(cfg.Rules) > 0 {
		return cfg.Rules
	}

	xRules := []string{"allow **/*"}
	for _, str := range cfg.Excludes {
		xRules = append(xRules, "deny "+str)
	}
	return xRules
}

func convertRulesToStruct(rulesStrings []string) []rule {
	rules := []rule{}
	for _, str := range rulesStrings {
		if strings.HasPrefix(str, "allow ") {
			rules = append(rules, rule{Allow: true, Mask: strings.TrimPrefix(str, "allow ")})
		} else {
			rules = append(rules, rule{Allow: false, Mask: strings.TrimPrefix(str, "deny ")})
		}
	}
	return rules
}

func checkFileByRules(rules []rule, path string) bool {
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

func readConfig(filename string) (*config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func calcHashBytes(buf []byte) string {
	h := sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func calcHashFile(path string) string {
	buf, _ := os.ReadFile(path)
	return calcHashBytes(buf)
}

func loadConfig(fileName string) *config {
	if _, err := os.Stat(fileName); err != nil {
		return &config{Excludes: []string{}}
	}
	cfg, _ := readConfig(fileName)
	// if err != nil {
	// 	os.(1)
	// }
	return cfg
}

type fileInfo struct {
	fileName string
	hash     string
}

func calcHashFiles(files []fileInfo) string {
	s := ""
	for _, file := range files {
		s += file.fileName + " " + file.hash + "\n"
	}
	return calcHashBytes([]byte(s))
}

func sortFiles(files []fileInfo) []fileInfo {
	sort.Slice(files, func(i, j int) bool {
		if files[i].fileName == files[j].fileName {
			return files[i].hash < files[j].hash
		}
		return (files[i].fileName < files[j].fileName)
	})
	return files
}

func makeFileList(cfg *config, root string) []fileInfo {
	rules := convertRulesToStruct(convert(cfg))
	files := make([]fileInfo, 0)
	err := filepath.Walk(root, func(path0 string, info os.FileInfo, err error) error {
		path, _ := filepath.Rel(root, path0)
		if info.IsDir() || checkFileByRules(rules, path) {
			return nil
		}
		files = append(files, fileInfo{fileName: filepath.ToSlash(path), hash: calcHashFile(path0)})
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func getShortHash(hash string, isShort bool) string {
	if isShort {
		return hash[:8]
	}
	return hash
}

func getRoot(root string) string {
	if root == "" {
		return "."
	}
	return root
}

func process(cfgPath string, root string) ([]fileInfo, string) {
	rootPath := getRoot(root)
	fullCfgPath := cfgPath
	if fullCfgPath == "" {
		fullCfgPath = rootPath + "/.prj2hash.yaml"
	}
	files := sortFiles(makeFileList(loadConfig(fullCfgPath), rootPath))
	return files, calcHashFiles(files)
}

var gitHash = "development"
var p2hHash = ""

func main() {
	isShort := flag.Bool("short", false, "Show short variant of hash")
	isHelp := flag.Bool("help", false, "Show help")
	isDryRun := flag.Bool("dry-run", false, "Show file list")
	cfgPath := flag.String("cfg", "", "config file for project")
	isVersion := flag.Bool("version", false, "Show version of application")
	flag.Parse()

	if *isVersion {
		fmt.Println("Version:")
		fmt.Println("     git", gitHash)
		if p2hHash != "" {
			fmt.Println("     p2h", p2hHash)
		}
		return
	}

	if *isHelp {
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		return
	}

	files, hash := process(*cfgPath, flag.Arg(0))
	if *isDryRun {
		for _, file := range files {
			fmt.Println(" file", file.fileName, file.hash)
		}
		fmt.Println("total", getShortHash(hash, true), hash)
	} else {
		fmt.Println(getShortHash(hash, *isShort))
	}
}
