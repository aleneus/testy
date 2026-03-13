// Package testy provides the assert helpers for writing unit tests.
package testy

import (
	"encoding/hex"
	"strings"
	"testing"
)

// DecodeHex concatenates parts, removes spaces, and decodes the hex
// string. It fails the test if decoding fails.
func DecodeHex(t *testing.T, parts ...string) []byte {
	t.Helper()

	var builder strings.Builder

	for _, p := range parts {
		builder.WriteString(p)
	}

	raw := strings.ReplaceAll(builder.String(), " ", "")

	bytes, err := hex.DecodeString(raw)
	AssertNoErr(t, err)

	return bytes
}
