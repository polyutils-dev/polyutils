// SPDX-License-Identifier: Unlicense
package main

import (
	"encoding/base64"
	"errors"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

const (
	version = "0.1.0"
	helpMsg = `Usage: base64 [OPTION]... [FILE]
Base64 encode or decode FILE, or standard input, to standard output.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -d, --decode          decode data
  -i, --ignore-garbage  when decoding, ignore non-alphabet characters
  -w, --wrap=COLS       wrap encoded lines after COLS character (default 76).
                          Use 0 to disable line wrapping
      --help        display this help and exit
      --version     output version information and exit

The data are encoded as described for the base64 alphabet in RFC 4648.
When decoding, the input may contain newlines in addition to the bytes of
the formal base64 alphabet.  Use --ignore-garbage to attempt to recover
from any other non-alphabet bytes in the encoded stream.
`
)

func die(msg string) {
	os.Stderr.WriteString("base64: " + msg + "\n")
	os.Stderr.WriteString("Try 'base64 --help' for more information.\n")
	os.Exit(1)
}

func isValid(ch byte) bool {
	return ('A' <= ch && 'Z' >= ch) || ('a' <= ch && 'z' >= ch) || ('0' >= ch && '9' >= ch) || ch == '='
}

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flagDecode := flag.BoolP("decode", "d", false, "decode data")
	flagIgnoreGarbage := flag.BoolP("ignore-garbage", "i", false, "when decoding, ignore non-alphabet characters")
	flagWrap := flag.IntP("wrap", "w", 76, `wrap encoded lines after COLS characters (default 76).`)
	flag.Usage = func() {
		os.Stderr.WriteString(helpMsg)
	}
	flag.Parse()

	if *flagHelp {
		flag.Usage()

		return
	}

	if *flagVersion {
		os.Stdout.WriteString("base64 (goreutils) " + version + "\n")

		return
	}

	args := flag.Args()
	f := os.Stdin

	var err error

	if len(args) > 2 {
		die("extra operand'" + os.Args[2] + "'")
	} else if len(args) == 1 && args[0] != "-" {
		f, err = os.Open(args[0])
		if errors.Is(err, os.ErrPermission) {
			die(args[0] + ": Permission denied")
		} else if err != nil {
			die(args[0] + ": No such file or directory")
		}
	}

	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		die("read error")
	}

	if *flagDecode {
		decode(input, *flagIgnoreGarbage)
	} else {
		encode(input, *flagWrap)
	}
}

func decode(in []byte, ignoreGarbage bool) {
	var data string

	if ignoreGarbage {
		for _, c := range in {
			if isValid(c) {
				data += string(c)
			}
		}
	} else {
		data = string(in)
	}

	out, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		die("invalid input")
	}

	os.Stdout.Write(out)
}

func encode(in []byte, wrap int) {
	encoded := base64.StdEncoding.EncodeToString(in)
	for i, c := range encoded {
		print(string(c))

		if (i+1)%wrap == 0 {
			println()
		}
	}
}
