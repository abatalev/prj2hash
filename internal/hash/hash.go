package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
)

func GetShortHash(hash string, isShort bool) string {
	if isShort {
		return hash[:8]
	}
	return hash
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
