// SPDX-License-Identifier: Unlicense
package main

import (
	"os"
	"os/user"

	flag "github.com/spf13/pflag"
)

const version = "0.1.0"

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: whoami [OPTION]...
Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit
`)
	}
	flag.Parse()

	if *flagHelp {
		flag.Usage()

		return
	}

	if *flagVersion {
		os.Stdout.WriteString("whoami (goreutils) " + version + "\n")

		return
	}

	u, err := user.Current()
	if err != nil {
		os.Stderr.WriteString("whoami: failed to execute")
		os.Exit(1)
	}

	os.Stdout.WriteString(u.Username + "\n")
}
