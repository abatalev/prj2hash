package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"gopkg.in/yaml.v3"
)

type config struct {
	Excludes []string `yaml:"excludes"`
}

func readConfig(filename string) (*config, error) {
	buf, err := ioutil.ReadFile(filename)
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
	buf, _ := ioutil.ReadFile(path)
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

func excludeMask(cfg *config, path string) bool {
	for _, mask := range cfg.Excludes {
		matched, _ := regexp.MatchString(mask, path)
		if matched {
			return true
		}
	}
	return false
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
	files := make([]fileInfo, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || excludeMask(cfg, path) {
			return nil
		}
		files = append(files, fileInfo{fileName: path, hash: calcHashFile(path)})
		return nil
	})
	return files
}

func getShortHash(hash string, isShort bool) string {
	if isShort {
		return hash[:8]
	}
	return hash
}

func main() {
	isShort := flag.Bool("short", false, "Show short variant of hash")
	isHelp := flag.Bool("help", false, "Show help")
	isDryRun := flag.Bool("dry-run", false, "Show file list")
	flag.Parse()
	if *isHelp {
		fmt.Println()
		flag.PrintDefaults()
		fmt.Println()
		return
	}

	root := flag.Arg(0)
	if root == "" {
		root = "."
	}

	files := sortFiles(makeFileList(loadConfig(".prj2hash.yaml"), root))
	hash := calcHashFiles(files)
	if *isDryRun {
		for _, file := range files {
			fmt.Println(" file", file.fileName, file.hash)
		}
		fmt.Println("total", getShortHash(hash, true), hash)
	} else {
		fmt.Println(getShortHash(hash, *isShort))
	}
}
