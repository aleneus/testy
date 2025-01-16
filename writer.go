// Package testy provides the assert helpers for writing unit tests.
package testy

// Writer is the very simple text writer. It can be used for testing
// log messages for example.
type Writer struct {
	text string
}

// Write adds data to writer.
func (w *Writer) Write(data []byte) (int, error) {
	w.text += string(data)

	return len(data), nil
}

// Pop returns all content and clears writer.
func (w *Writer) Pop() string {
	res := w.text
	w.text = ""

	return res
}
