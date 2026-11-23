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
	cmd := exec.Command("go", "run", ".", "-r", "Drive")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("hello -r Drive: %v\n%s", err, out)
	}
	if got := string(out); !strings.Contains(got, "evirD") {
		t.Errorf("hello -r Drive output %q, want it to contain %q", got, "evirD")
	}
}

func TestUpperFlag(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-u", "Drive")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("hello -u Drive: %v\n%s", err, out)
	}
	if got := string(out); !strings.Contains(got, "HELLO, DRIVE!") {
		t.Errorf("hello -u Drive output %q, want it to contain %q", got, "HELLO, DRIVE!")
	}
}

func TestUpperReverseFlag(t *testing.T) {
	cmd := exec.Command("go", "run", ".", "-u", "-r", "Drive")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("hello -u -r Drive: %v\n%s", err, out)
	}
	if got := string(out); !strings.Contains(got, "HELLO, EVIRD!") {
		t.Errorf("hello -u -r Drive output %q, want it to contain %q", got, "HELLO, EVIRD!")
	}
}
