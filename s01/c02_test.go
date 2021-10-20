package s01

import (
	"encoding/hex"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	data := "1c0111001f010100061a024b53535009181c"
	d, err := hex.DecodeString(data)
	if err != nil {
		t.Fatal("failed to decode")
	}

	key := "686974207468652062756c6c277320657965"
	k, err := hex.DecodeString(key)
	if err != nil {
		t.Fatal("failed to decode")
	}

	b := []byte{}
	l := len(k)
	for i := range d {
		b = append(b, d[i]^k[i%l])
	}

	exp := "746865206b696420646f6e277420706c6179"
	act := hex.EncodeToString(b)
	if exp != act {
		t.Fatalf("expected %q but got %q", exp, act)
	}
}
