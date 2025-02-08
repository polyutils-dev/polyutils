# polyutils

The GNU coreutils, rewritten in every language.

[![License: Unlicense](https://img.shields.io/badge/license-unlicense-blue.svg)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/polyutils/polyutils)](https://goreportcard.com/report/codeberg.org/polyutils/polyutils)

## Roadmap

* [ ] `chcon`
* [ ] `chgrp`
* [ ] `chown`
* [ ] `chmod`
* [ ] `cp`
* [ ] `dd`
* [ ] `df`
* [ ] `dir`
* [ ] `dircolors`
* [ ] `install`
* [ ] `ln`
* [ ] `ls`
* [ ] `mkdir`
* [ ] `mkfifo`
* [ ] `mknod`
* [ ] `mktemp`
* [ ] `mv`
* [ ] `realpath`
* [ ] `rm`
* [ ] `rmdir`
* [ ] `shred`
* [ ] `sync`
* [ ] `touch`
* [ ] `truncate`
* [ ] `vdir`
* [ ] `b2sum`
* [x] `base32`
* [x] `base64`
* [ ] `basenc`
* [ ] `cat`
* [ ] `cksum`
* [ ] `comm`
* [ ] `csplit`
* [ ] `cut`
* [ ] `expand`
* [ ] `fmt`
* [ ] `fold`
* [ ] `head`
* [ ] `join`
* [ ] `md5sum`
* [ ] `nl`
* [ ] `numfmt`
* [ ] `od`
* [ ] `paste`
* [ ] `ptx`
* [ ] `pr`
* [ ] `sha1sum`
* [ ] `sha224sum`
* [ ] `sha256sum`
* [ ] `sha384sum`
* [ ] `sha512sum`
* [ ] `shuf`
* [ ] `sort`
* [ ] `split`
* [ ] `sum`
* [ ] `tac`
* [ ] `tail`
* [ ] `tr`
* [ ] `tsort`
* [ ] `unexpand`
* [ ] `uniq`
* [ ] `wc`
* [ ] `arch`
* [ ] `basename`
* [ ] `chroot`
* [ ] `date`
* [ ] `dirname`
* [ ] `du`
* [x] `echo`
* [ ] `env`
* [ ] `expr`
* [ ] `factor`
* [x] `false`
* [ ] `groups`
* [ ] `hostid`
* [ ] `id`
* [ ] `link`
* [ ] `logname`
* [ ] `nice`
* [ ] `nohup`
* [ ] `nproc`
* [ ] `pathchk`
* [ ] `pinky`
* [ ] `printenv`
* [ ] `printf`
* [ ] `pwd`
* [ ] `readlink`
* [ ] `runcon`
* [ ] `seq`
* [x] `sleep`
* [ ] `stat`
* [ ] `stdbuf`
* [ ] `stty`
* [ ] `tee`
* [ ] `test`
* [ ] `timeout`
* [x] `true`
* [ ] `tty`
* [ ] `uname`
* [ ] `unlink`
* [ ] `uptime`
* [ ] `users`
* [ ] `who`
* [x] `whoami`
* [x] `yes`
  * So some info: when built with `go`, it's extremely fast, but the binary is
    bigger. However, with `tinygo`, it's about half as fast, but the binary is
    much smaller. You should choose which is more important to you.
* [ ] `[`

## License

This project is licensed under the Unlicense.
