package hash

func GetShortHash(hash string, isShort bool) string {
	if isShort {
		return hash[:8]
	}
	return hash
}
