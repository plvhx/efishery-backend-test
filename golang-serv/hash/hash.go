package hash

import (
	"fmt"
	"io"
	"crypto/sha1"
)

func Calculate(data string) string {
	hash := sha1.New()
	io.WriteString(hash, data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
