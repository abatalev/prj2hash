package main

import (
	"flag"
	"fmt"

	"abatalev.com/prj2hash/internal/config"
	"abatalev.com/prj2hash/internal/files"
	"abatalev.com/prj2hash/internal/hash"
)

var gitHash = "development"
var p2hHash = ""

func getRoot(root string) string {
	if root == "" {
		return "."
	}
	return root
}

func process(cfgPath string, root string) ([]files.FileInfo, string) {
	rootPath := getRoot(root)
	fullCfgPath := cfgPath
	if fullCfgPath == "" {
		fullCfgPath = rootPath + "/.prj2hash.yaml"
	}
	fileList := files.SortFiles(files.MakeFileList(config.LoadConfig(fullCfgPath), rootPath))
	return fileList, files.CalcHashFiles(fileList)
}

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

	files, prjHash := process(*cfgPath, flag.Arg(0))
	if *isDryRun {
		for _, file := range files {
			fmt.Println(" file", file.FileName, file.Hash)
		}
		fmt.Println("total", hash.GetShortHash(prjHash, true), prjHash)
	} else {
		fmt.Println(hash.GetShortHash(prjHash, *isShort))
	}
}
