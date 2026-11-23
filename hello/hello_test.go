// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestReverseFlag(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "hello")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}

	got, err := exec.Command(bin, "-r", "Drive").CombinedOutput()
	if err != nil {
		t.Fatalf("hello -r Drive: %v\n%s", err, got)
	}

	want := "olleH, evirD!\n"
	if string(got) != want {
		t.Errorf("hello -r Drive == %q, want %q", got, want)
	}

	if !strings.Contains(string(got), "evirD") {
		t.Errorf("hello -r Drive == %q, want reversed greeting containing %q", got, "evirD")
	}
}
