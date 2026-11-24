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
	if got := string(out); !strings.Contains(got, "evirD") {
		t.Errorf("hello -r Drive output %q, want it to contain %q", got, "evirD")
	}
}
