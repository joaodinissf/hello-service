// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainGreeting(t *testing.T) {
	dir := t.TempDir()
	bin := filepath.Join(dir, "hello")
	build := exec.Command("go", "build", "-o", bin, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}

	for _, c := range []struct {
		name string
		args []string
		want string
	}{
		{"reverse", []string{"-r", "Drive"}, "olleH, evirD!\n"},
		{"plain", []string{"Drive"}, "Hello, Drive!\n"},
		{"upper", []string{"-u", "Drive"}, "HELLO, DRIVE!\n"},
		{"upper-reverse", []string{"-u", "-r", "Drive"}, "OLLEH, EVIRD!\n"},
	} {
		t.Run(c.name, func(t *testing.T) {
			cmd := exec.Command(bin, c.args...)
			cmd.Env = os.Environ()
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("%s: %v\n%s", bin, err, out)
			}
			if got := string(out); got != c.want {
				t.Errorf("hello %v == %q, want %q", c.args, got, c.want)
			}
		})
	}
}
