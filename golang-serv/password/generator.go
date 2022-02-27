package password

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowercase  = "abcdefhijklmnopqrstuvwxyz"
	uppercase  = "ABCDEFHIJKLMNOPQRSTUVWXYZ"
	digits     = "0123456789"
	allCharset = lowercase + uppercase + digits
)

func Generate(length int) string {
	rand.Seed(time.Now().Unix())

	var password strings.Builder

	clen := len(allCharset)

	for i := 0; i < length; i++ {
		index := rand.Intn(clen)
		password.WriteString(string(allCharset[index]))
	}

	prune := []rune(password.String())
	lprune := len(prune)

	rand.Shuffle(lprune, func(i, j int) {
		prune[i], prune[j] = prune[j], prune[i]
	})

	return string(prune)
}
