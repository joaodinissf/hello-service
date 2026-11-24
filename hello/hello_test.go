// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestReverseGreeting(t *testing.T) {
	bin := t.TempDir() + "/hello"
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}

	out, err := exec.Command(bin, "-r", "Drive").CombinedOutput()
	if err != nil {
		t.Fatalf("hello -r Drive: %v\n%s", err, out)
	}

	want := "olleH, evirD!\n"
	if got := string(out); got != want {
		t.Errorf("hello -r Drive == %q, want %q", got, want)
	}
	if !strings.Contains(string(out), "evirD") {
		t.Errorf("hello -r Drive == %q, want name greeted in reverse", out)
	}
}
