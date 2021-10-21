package s01

import (
	"encoding/hex"
	"strings"
	"testing"
)

func score(b []byte) int {
	return len(strings.Split(string(b), " "))
}

func TestSingleByteXORCipher(t *testing.T) {
	var err error

	// the string `data` should produce some English plantext when xor'd with
	// one of the bytes from `keys`
	data := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	keys := []byte(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var db []byte
	if db, err = hex.DecodeString(data); err != nil {
		t.Fatal(err)
	}

	best, guess := 0, []byte{}
	for _, k := range keys {
		pb := []byte{}
		for i := range db {
			pb = append(pb, db[i]^k)
		}
		if s := score(pb); s > best {
			best, guess = s, pb
		}
	}

	t.Logf("db = %s", string(db))
	t.Logf("best = %d", best)
	t.Logf("guess = %s", string(guess))
}
