// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"testing"
)

func TestGreet(t *testing.T) {
	for _, c := range []struct {
		name        string
		greeting    string
		arg         string
		reverseFlag bool
		upperFlag   bool
		want        string
	}{
		{"default", "Hello", "world", false, false, "Hello, world!"},
		{"named", "Hello", "Drive", false, false, "Hello, Drive!"},
		{"upper", "Hello", "Drive", false, true, "HELLO, DRIVE!"},
		{"upper world", "Hello", "world", false, true, "HELLO, WORLD!"},
		{"reverse", "Hello", "Drive", true, false, "olleH, evirD!"},
		{"upper reverse", "Hello", "Drive", true, true, "OLLEH, EVIRD!"},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := greet(c.greeting, c.arg, c.reverseFlag, c.upperFlag)
			if got != c.want {
				t.Errorf("greet(%q, %q, %v, %v) == %q, want %q",
					c.greeting, c.arg, c.reverseFlag, c.upperFlag, got, c.want)
			}
		})
	}
}

func TestUpperFlagAliases(t *testing.T) {
	for _, c := range []struct {
		name string
		arg  string
	}{
		{"short", "-u"},
		{"long", "--upper"},
	} {
		t.Run(c.name, func(t *testing.T) {
			var upper bool
			fs := flag.NewFlagSet("hello", flag.ContinueOnError)
			fs.BoolVar(&upper, "u", false, "")
			fs.BoolVar(&upper, "upper", false, "")
			if err := fs.Parse([]string{c.arg, "Drive"}); err != nil {
				t.Fatalf("Parse(%q): %v", c.arg, err)
			}
			if !upper {
				t.Errorf("%s did not set upper", c.arg)
			}
			if got, want := greet("Hello", fs.Arg(0), false, upper), "HELLO, DRIVE!"; got != want {
				t.Errorf("greet with %s == %q, want %q", c.arg, got, want)
			}
		})
	}
}
