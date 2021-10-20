package s01

import (
	"encoding/hex"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	// the string `data` should produce the string `exp` when xor'd with the
	// string `key`
	data := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	exp := "746865206b696420646f6e277420706c6179"

	var err error
	var db, kb []byte

	if db, err = hex.DecodeString(data); err != nil {
		t.Fatal("failed to decode")
	}
	if kb, err = hex.DecodeString(key); err != nil {
		t.Fatal("failed to decode")
	}

	eb, l := []byte{}, len(kb)
	for i := range db {
		eb = append(eb, db[i]^kb[i%l])
	}

	act := hex.EncodeToString(eb)
	if exp != act {
		t.Fatalf("expected %q but got %q", exp, act)
	}
	
	t.Logf("db = %s", string(db))
	t.Logf("kb = %s", string(kb))
	t.Logf("eb = %s", string(eb))
}
