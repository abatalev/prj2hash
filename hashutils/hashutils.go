package hashutils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"sort"
)

type FileInfo struct {
	FileName string
	Hash     string
}

func CalcHashBytes(buf []byte) string {
	h := sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil))
}

func CalcHashFile(path string) string {
	buf, _ := os.ReadFile(path)
	return CalcHashBytes(buf)
}

func CalcHashFiles(fileList []FileInfo) string {
	s := ""
	for _, file := range fileList {
		s += file.FileName + " " + file.Hash + "\n"
	}
	return CalcHashBytes([]byte(s))
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
