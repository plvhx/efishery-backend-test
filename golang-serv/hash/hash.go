package hash

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Calculate(data string) string {
	hash := sha1.New()
	io.WriteString(hash, data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
