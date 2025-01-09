// SPDX-License-Identifier: Unlicense
package main

import (
	"os"

	flag "github.com/spf13/pflag"
)

var (
	authors = []string{"Zia M"}
	version = "0.1.0"
)

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: false [ignored command line arguments]
  or:  false OPTION
Exit with a status code indicating success.

      --help        display this help and exit
      --version     output version information and exit

NOTE: your shell may have its own version of false, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.
`)
	}
	flag.Parse()

	if *flagHelp {
		flag.Usage()
	}

	if *flagVersion {
		os.Stdout.WriteString("false (goreutils) 0.1.0\n")
	}

	os.Exit(1)
}
