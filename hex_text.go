package testy

import (
	"testing"
)

func TestDecodeHex_BmpHeader(t *testing.T) {
	t.Parallel()

	t.Run("BMP header", func(t *testing.T) {
		t.Parallel()

		data := DecodeHex(t,
			"42 4D",       // Signature: "BM"
			"46 00 00 00", // File size: 70 bytes
			"00 00",       // Reserved 1
			"00 00",       // Reserved 2
			"36 00 00 00", // Data offset: 54 bytes
		)

		AssertEqual(t, len(data), 13)
	})

	t.Run("UDP header", func(t *testing.T) {
		t.Parallel()

		data := DecodeHex(t,
			"04 D2", // Source Port: 1234
			"00 50", // Destination Port: 80 (HTTP)
			"00 0C", // Length: 12 bytes (8 header + 4 data)
			"2E D0", // Checksum
		)

		AssertEqual(t, len(data), 8)
	})
}
