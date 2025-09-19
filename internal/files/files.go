package files

import (
	"os"
	"path/filepath"

	libhash "github.com/abatalev/prj2hash/hashutils"
	"github.com/abatalev/prj2hash/internal/config"
	"github.com/abatalev/prj2hash/internal/rules"
)

func MakeFileList(cfg *config.Config, root string) []libhash.FileInfo {
	ruleList := rules.ConvertRulesToStruct(rules.Convert(cfg))
	files := make([]libhash.FileInfo, 0)
	err := filepath.Walk(root, func(path0 string, info os.FileInfo, _ error) error {
		path, _ := filepath.Rel(root, path0)
		if info.IsDir() || rules.CheckFileByRules(ruleList, path) {
			return nil
		}
		files = append(files, libhash.FileInfo{FileName: filepath.ToSlash(path), Hash: libhash.CalcHashFile(path0)})
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
