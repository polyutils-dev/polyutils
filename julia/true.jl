for arg in ARGS
    if arg == "--help"
        println("""Usage: true [ignored command line arguments]
  or:  true OPTION
Exit with a status code indicating success.

      --help        display this help and exit
      --version     output version information and exit

NOTE: your shell may have its own version of true, which usually supersedes
the version described here.  Please refer to your shell's documentation
for details about the options it supports.""")
        exit(0)
    end

    if arg == "--version"
        println("true (julia polyutils) 0.1.0")
        exit(0)
    end
end

exit(0)
