#!/usr/bin/env python3
import argparse
import os
import sys

import lib

def main() -> int:
    parser = argparse.ArgumentParser(prog="whoami",
                                     add_help=False)

    parser.add_argument("--version", action="store_true")
    parser.add_argument("--help", action="store_true")

    args = parser.parse_args()

    if args.help:
        print("""Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit""")
        return 0

    if args.version:
        print("whoami (python polyutils) " + lib.__version__)
        return 0

    print(os.getlogin())
    
    return 0

if __name__ == "__main__":
    sys.exit(main())
