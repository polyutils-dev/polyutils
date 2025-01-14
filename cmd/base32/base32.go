package main

import (
	"encoding/base32"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	authors = []string{"Zia M"}
	version = "0.1.0"
)

func die(msg string) {
	os.Stderr.WriteString("base32: " + msg + "\n")
	os.Stderr.WriteString("Try 'base32 --help' for more information.\n")
	os.Exit(1)
}

func isValid(ch byte) bool {
	return ('A' <= ch && 'Z' >= ch) || ('2' <= ch && '7' >= ch)
}

func main() {
	flagHelp := flag.Bool("help", false, "display this help and exit")
	flagVersion := flag.Bool("version", false, "output version information and exit")
	flagDecode := flag.BoolP("decode", "d", false, "decode data")
	flagIgnoreGarbage := flag.BoolP("ignore-garbage", "i", false, "when decoding, ignore non-alphabet characters")
	flagWrap := flag.IntP("wrap", "w", 76, `wrap encoded lines after COLS characters (default 76).`)
	flag.Usage = func() {
		os.Stderr.WriteString(`Usage: base32 [OPTION]... [FILE]
Base32 encode or decode FILE, or standard input, to standard output.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -d, --decode          decode data
  -i, --ignore-garbage  when decoding, ignore non-alphabet characters
  -w, --wrap=COLS       wrap encoded lines after COLS character (default 76).
                          Use 0 to disable line wrapping
      --help        display this help and exit
      --version     output version information and exit

The data are encoded as described for the base32 alphabet in RFC 4328.
When decoding, the input may contain newlines in addition to the bytes of
the formal base32 alphabet.  Use --ignore-garbage to attempt to recover
from any other non-alphabet bytes in the encoded stream.
`)
	}
	flag.Parse()

	if *flagHelp {
		flag.Usage()
		return
	}
	if *flagVersion {
		os.Stdout.WriteString("base32 (goreutils) 0.1.0")
		return
	}

	args := flag.Args()
	f := os.Stdin
	var err error

	if len(args) > 2 {
		die("extra operand'" + os.Args[2] + "'")
	}
	if len(args) == 1 && args[0] != "-" {
		f, err = os.Open(args[0])
		if err == os.ErrPermission {
			die(args[0] + ": Permission denied")
		}
		if err != nil {
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
		data = string(in[:])
	}

	out, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		die("invalid input")
	}
	os.Stdout.Write(out)
	os.Stdout.Write([]byte{'\n'})
}

func encode(in []byte, wrap int) {
	encoded := base32.StdEncoding.EncodeToString(in)
	for i := 0; i < len(encoded); i++ {
		print(string(encoded[i]))
		if (i+1)%wrap == 0 {
			println()
		}
	}
}
