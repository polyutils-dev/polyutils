// SPDX-License-Identifier: Unlicense
package main

import (
	"bytes"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

const version = "0.1.0"

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: yes [STRING]...
  or:  yes OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

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
		os.Stdout.WriteString("yes (goreutils) " + version + "\n")

		return
	}

	str := strings.Join(flag.Args(), "")
	if str == "" {
		str = "y"
	}

	value := bytes.Repeat([]byte(str+"\n"), 100_000_000)

	for {
		os.Stdout.Write(value)
	}
}
