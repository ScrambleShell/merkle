package merkle

import (
	"crypto/sha256"
	"encoding/base64"
)

// Hash generates a merkle root of input elements SHA256 hash
func Hash(elements []string) string {
	// base cases
	if len(elements) == 0 {
		return ""
	} else if len(elements) == 1 {
		return elements[0]
	}

	half := len(elements) / 2
	a := Hash(elements[:half])
	b := Hash(elements[half:])
	return hashTwoHash(a, b)
}

func hashTwoHash(a, b string) string {
	combined := a + b
	hash := sha256.New()
	hash.Write([]byte(combined))
	sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return sha
}
