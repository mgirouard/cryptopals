package c01

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestHex2Base64(t *testing.T) {
	h := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b, err := hex.DecodeString(h)
	if err != nil {
		t.Fatal("failed to decode")
	}
	exp := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	act := base64.StdEncoding.EncodeToString(b)
	if exp != act {
		t.Fatalf("expected %q but got %q", exp, act)
	}
}
