// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestReverseFlag(t *testing.T) {
	out, err := exec.Command("go", "run", ".", "-r", "Drive").CombinedOutput()
	if err != nil {
		t.Fatalf("go run . -r Drive: %v\n%s", err, out)
	}
	got := strings.TrimSpace(string(out))
	want := "olleH, evirD!"
	if got != want {
		t.Errorf("go run . -r Drive == %q, want %q", got, want)
	}
}

func TestUpperFlag(t *testing.T) {
	out, err := exec.Command("go", "run", ".", "-u", "Drive").CombinedOutput()
	if err != nil {
		t.Fatalf("go run . -u Drive: %v\n%s", err, out)
	}
	got := strings.TrimSpace(string(out))
	want := "HELLO, DRIVE!"
	if got != want {
		t.Errorf("go run . -u Drive == %q, want %q", got, want)
	}
}

func TestUpperReverseFlag(t *testing.T) {
	out, err := exec.Command("go", "run", ".", "-u", "-r", "Drive").CombinedOutput()
	if err != nil {
		t.Fatalf("go run . -u -r Drive: %v\n%s", err, out)
	}
	got := strings.TrimSpace(string(out))
	want := "OLLEH, EVIRD!"
	if got != want {
		t.Errorf("go run . -u -r Drive == %q, want %q", got, want)
	}
}
