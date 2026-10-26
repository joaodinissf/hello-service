// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	if os.Getenv("HELLO_BE_CHILD") == "1" {
		main()
		os.Exit(0)
	}
	os.Exit(m.Run())
}

func TestFlags(t *testing.T) {
	for _, c := range []struct {
		name string
		args []string
		want string
	}{
		{"upper", []string{"-u", "Drive"}, "HELLO, DRIVE!"},
		{"upper reverse", []string{"-u", "-r", "Drive"}, "OLLEH, EVIRD!"},
		{"upper greeting", []string{"-u", "-g", "hi", "Drive"}, "HI, DRIVE!"},
	} {
		t.Run(c.name, func(t *testing.T) {
			cmd := exec.Command(os.Args[0], c.args...)
			cmd.Env = append(os.Environ(), "HELLO_BE_CHILD=1")
			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("running %v: %v", c.args, err)
			}
			got := strings.TrimSpace(string(out))
			if got != c.want {
				t.Errorf("hello %s == %q, want %q", strings.Join(c.args, " "), got, c.want)
			}
		})
	}
}
