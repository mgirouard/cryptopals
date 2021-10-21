package s01

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestHex2Base64(t *testing.T) {
	var err error

	// the hex string `data` should produce the string `exp` when base64
	// encoded
	data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	exp := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	var db []byte
	if db, err = hex.DecodeString(data); err != nil {
		t.Fatal(err)
	}

	act := base64.StdEncoding.EncodeToString(db)
	if exp != act {
		t.Fatalf("expected %q but got %q", exp, act)
	}

	t.Logf("db = %s", string(db))
}
