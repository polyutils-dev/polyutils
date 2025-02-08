for arg in ARGS
    if arg == "--help"
        println("""Print the user name associated with the current effective user ID.
Same as id -un.

      --help        display this help and exit
      --version     output version information and exit""")

        exit(0)
    end

    if arg == "--version"
        println("whoami (julia polyutils) 0.1.0")
        exit(0)
    end

    #= FIXME: This command ignores extra parameters and flags, whereas other
    # implementations do not. =#
end

println(Sys.username())
