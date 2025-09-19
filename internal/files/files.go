package files

import (
	"os"
	"path/filepath"
	"sort"

	"abatalev.com/prj2hash/internal/config"
	"abatalev.com/prj2hash/internal/hash"
	"abatalev.com/prj2hash/internal/rules"
)

type FileInfo struct {
	FileName string
	Hash     string
}

func MakeFileList(cfg *config.Config, root string) []FileInfo {
	ruleList := rules.ConvertRulesToStruct(rules.Convert(cfg))
	files := make([]FileInfo, 0)
	err := filepath.Walk(root, func(path0 string, info os.FileInfo, _ error) error {
		path, _ := filepath.Rel(root, path0)
		if info.IsDir() || rules.CheckFileByRules(ruleList, path) {
			return nil
		}
		files = append(files, FileInfo{FileName: filepath.ToSlash(path), Hash: hash.CalcHashFile(path0)})
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func SortFiles(files []FileInfo) []FileInfo {
	sort.Slice(files, func(i, j int) bool {
		if files[i].FileName == files[j].FileName {
			return files[i].Hash < files[j].Hash
		}
		return (files[i].FileName < files[j].FileName)
	})
	return files
}

func CalcHashFiles(fileList []FileInfo) string {
	s := ""
	for _, file := range fileList {
		s += file.FileName + " " + file.Hash + "\n"
	}
	return hash.CalcHashBytes([]byte(s))
}
