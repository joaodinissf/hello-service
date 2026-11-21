// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"testing"
)

func TestGreetUpper(t *testing.T) {
	if got, want := greet([]string{"Drive"}, "Hello", false, true), "HELLO, DRIVE!\n"; got != want {
		t.Errorf("greet(-u Drive) = %q, want %q", got, want)
	}
}

func TestGreetUpperReverse(t *testing.T) {
	if got, want := greet([]string{"Drive"}, "Hello", true, true), "OLLEH, EVIRD!\n"; got != want {
		t.Errorf("greet(-u -r Drive) = %q, want %q", got, want)
	}
}

func TestUpperFlags(t *testing.T) {
	for _, c := range []struct {
		name string
		args []string
		want bool
	}{
		{"short", []string{"-u", "Drive"}, true},
		{"long", []string{"--upper", "Drive"}, true},
		{"none", []string{"Drive"}, false},
	} {
		t.Run(c.name, func(t *testing.T) {
			fs := flag.NewFlagSet("hello", flag.ContinueOnError)
			upperFlag := fs.Bool("u", false, "")
			upperLong := fs.Bool("upper", false, "")
			if err := fs.Parse(c.args); err != nil {
				t.Fatalf("parse %v: %v", c.args, err)
			}
			if got := *upperFlag || *upperLong; got != c.want {
				t.Errorf("%v upper = %v, want %v", c.args, got, c.want)
			}
		})
	}
}
