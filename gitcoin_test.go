package gitcoin_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/benbjohnson/gitcoin"
)

func TestTarget_Check(t *testing.T) {
	target := gitcoin.NewTarget()
	if !target.Check("foo") {
		t.Fatalf("expected check to succeed")
	}
	if !bytes.Equal(target.Value, gitcoin.Hash("foo")) {
		t.Fatalf("unexpected new target: %x", target.Value)
	}
}

// Ensure the hash returns a SHA1 hash.
func TestHash(t *testing.T) {
	v := gitcoin.Hash("foo")
	if s := fmt.Sprintf("%x", v); s != "0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33" {
		t.Fatalf("unexpected hash: %s", s)
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gitcoin.Hash("foo")
	}
}
