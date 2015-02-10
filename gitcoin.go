package gitcoin

import (
	"bytes"
	"crypto/sha1"
	"sync"
)

var MaxValue = []byte{
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
}

type Target struct {
	mu    sync.Mutex
	Value []byte
}

// NewTarget returns a new instance of Target with the highest value possible.
func NewTarget() *Target {
	return &Target{Value: MaxValue}
}

// Check returns true if the string hashes to less than the target.
// The target is changed to the new hash if successful.
func (t *Target) Check(s string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Generate hash for string.
	v := Hash(s)

	// If the hash is not less than the current target then return false.
	if bytes.Compare(v, t.Value) != -1 {
		return false
	}

	// If it's less than target then update target and return true.
	t.Value = v
	return true
}

// Hash computes a SHA1 hash on the string.
func Hash(s string) []byte {
	value := sha1.Sum([]byte(s))
	return value[:]
}
