// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestGreet(t *testing.T) {
	for _, c := range []struct {
		name     string
		greeting string
		subject  string
		reverse  bool
		upper    bool
		want     string
	}{
		{"plain", "Hello", "Drive", false, false, "Hello, Drive!"},
		{"upper", "Hello", "Drive", false, true, "HELLO, DRIVE!"},
		{"reverse", "Hello", "Drive", true, false, "olleH, evirD!"},
		{"upper and reverse", "Hello", "Drive", true, true, "OLLEH, EVIRD!"},
	} {
		got := greet(c.greeting, c.subject, c.reverse, c.upper)
		if got != c.want {
			t.Errorf("%s: greet(%q, %q, %v, %v) == %q, want %q", c.name, c.greeting, c.subject, c.reverse, c.upper, got, c.want)
		}
	}
}
