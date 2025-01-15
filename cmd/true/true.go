// SPDX-License-Identifier: Unlicense
package main

import (
	"os"

	flag "github.com/spf13/pflag"
)

const version = "0.1.0"

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: true [ignored command line arguments]
  or:  true OPTION
Exit with a status code indicating success.

      --help        display this help and exit
      --version     output version information and exit

NOTE: your shell may have its own version of true, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.
`)
	}
	flag.Parse()

	if *flagHelp {
		flag.Usage()

		return
	}

	if *flagVersion {
		os.Stdout.WriteString("true (goreutils) " + version + "\n")
	}
}
