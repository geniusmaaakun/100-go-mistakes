package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCopySourceToDest(t *testing.T) {
	const input = "foo"

	// strings.NewReader() は io.Reader を実装している
	source := strings.NewReader(input)
	// bytes.Buffer は io.Writer を実装している
	dest := bytes.NewBuffer(make([]byte, 0))

	err := copySourceToDest(source, dest)
	if err != nil {
		t.FailNow()
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected: %s, got: %s", input, got)
	}
}
