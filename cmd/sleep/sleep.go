// SPDX-License-Identifier: Unlicense
package main

import (
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

func die(msg string) {
	os.Stderr.WriteString("sleep: " + msg + "\n")
	os.Stderr.WriteString("Try 'sleep --help' for more information.\n")
	os.Exit(1)
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: sleep NUMBER[SUFFIX]...
  or:  sleep OPTION
Pause for NUMBER seconds.  SUFFIX may be 's' for seconds (the default),
'm' for minutes, 'h' for hours or 'd' for days.  NUMBER need not be an
integer.  Given two or more arguments, pause for the amount of time
specified by the sum of their values.

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
		os.Stdout.WriteString("sleep (goreutils) 0.1.0\n")
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		die("missing operand")
	}

	durations := make([]time.Duration, len(args))
	var err error

	for i, arg := range args {
		if isNumber(arg[len(arg)-1]) {
			// The default wait time is seconds. If no unit is passed, then adding an
			// `s` at the end will make time.ParseDuration parse it as seconds.
			arg += "s"
		}

		durations[i], err = time.ParseDuration(arg)
		if err != nil {
			die("invalid time interval ‘" + arg + "’")
		}
	}

	for _, d := range durations {
		time.Sleep(d)
	}
}
