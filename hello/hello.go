// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a hello, world program, demonstrating
// how to write a simple command-line program.
//
// Usage:
//
//	hello [options] [name]
//
// The options are:
//
//	-g greeting
//		Greet with the given greeting, instead of "Hello".
//
//	-r
//		Greet in reverse.
//
//	-u, --upper
//		Greet in upper case.
//
// By default, hello greets the world.
// If a name is specified, hello greets that name instead.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/example/hello/reverse"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: hello [options] [name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	greeting    = flag.String("g", "Hello", "Greet with `greeting`")
	reverseFlag = flag.Bool("r", false, "Greet in reverse")
	upperFlag   bool
)

func init() {
	flag.BoolVar(&upperFlag, "u", false, "Greet in upper case")
	flag.BoolVar(&upperFlag, "upper", false, "Greet in upper case")
}

func main() {
	// Configure logging for a command-line program.
	log.SetFlags(0)
	log.SetPrefix("hello: ")

	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments.
	name := "world"
	args := flag.Args()
	if len(args) >= 2 {
		usage()
	}
	if len(args) >= 1 {
		name = args[0]
	}
	if name == "" { // hello '' is an error
		log.Fatalf("invalid name %q", name)
	}

	// Run actual logic.
	fmt.Println(greet(*greeting, name, *reverseFlag, upperFlag))
}

func greet(greeting, name string, reverseFlag, upperFlag bool) string {
	if reverseFlag {
		greeting = reverse.String(greeting)
		name = reverse.String(name)
	}
	out := fmt.Sprintf("%s, %s!", greeting, name)
	if upperFlag {
		out = strings.ToUpper(out)
	}
	return out
}
