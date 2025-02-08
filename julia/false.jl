for arg in ARGS
    if arg == "--help"
        println("""Usage: false [ignored command line arguments]
  or:  false OPTION
Exit with a status code indicating failure.

      --help        display this help and exit
      --version     output version information and exit

NOTE: your shell may have its own version of false, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.""")
        exit(1)
    end

    if arg == "--version"
        println("false (julia polyutils) 0.1.0")
        exit(1)
    end
end

exit(1)
