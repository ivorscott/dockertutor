package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestPrompt(t *testing.T) {
	want := "docker run hello-world"
	reader := bufio.NewReader(strings.NewReader(want))

	got, err := prompt(reader)
	if err != io.EOF {
		t.Errorf("Prompt failed and sent: %s", err)
	}

	if !bytes.Equal([]byte(got), []byte(want)) {
		t.Fatalf("Expected %s, got %s instead", want, got)
	}
}
