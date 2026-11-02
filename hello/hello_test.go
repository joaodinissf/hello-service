// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestGreeting(t *testing.T) {
	for _, c := range []struct {
		greeting, name string
		upper, reverse bool
		want           string
	}{
		{"Hello", "world", false, false, "Hello, world!"},
		{"Hello", "Drive", true, false, "HELLO, DRIVE!"},
		{"Hello", "Drive", true, true, "OLLEH, EVIRD!"},
		{"Hello", "Drive", false, true, "olleH, evirD!"},
		{"Hello", "", false, false, "Hello, !"},
		{"Hello", "Zoë", false, false, "Hello, Zoë!"},
		{"Hello", "Zoë", false, true, "olleH, ëoZ!"},
	} {
		got := greet(c.greeting, c.name, c.upper, c.reverse)
		if got != c.want {
			t.Errorf("greet(%q, %q, %v, %v) == %q, want %q", c.greeting, c.name, c.upper, c.reverse, got, c.want)
		}
	}
}
