package cli

import (
	"bufio"
	"github.com/farinap5/SSHPKM/src/db"
	"os"
	"strings"
)

var CMDStatus string = "\033[1;31m*\033[0;0m"

func Start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var c string
		print("[" + CMDStatus + "]> ")
		c, _ = reader.ReadString('\n')
		handleCmd(c)
	}
}

func handleCmd(c string) {
	c = strings.Split(c, "\n")[0]
	cs := strings.Split(c, " ")

	if c == "help" || c == "h" {
		Help()
	} else if c == "exit" || c == "quit" || c == "e" {
		os.Exit(0)
	} else if cs[0] == "create" {
		create(cs)
	} else if cs[0] == "list" {
		list(cs)
	} else if cs[0] == "config" {
		config(cs)
	}
}

func create(cs []string) {
	if len(cs) < 3 {
		println("More arguments needed.")
		return
	}
	if cs[1] == "user" {
		db.DBCreateUser(cs[2])
	} else if cs[1] == "host" {
		db.DBCreateHost(cs[2])
	} else if cs[1] == "help" {
		HelpCreate()
	} else {
		HelpCreate()
	}
}

func list(cs []string) {
	if len(cs) < 2 {
		println("More arguments needed.")
		return
	}
	if cs[1] == "user" {
		db.DBListUser()
	} else if cs[1] == "host" {
		db.DBListHost()
	}
}

func config(cs []string) {
	if cs[1] == "help" {
		HelpConfig()
		return
	}
	if len(cs) < 3 {
		println("More arguments needed.")
		return
	}
	if cs[1] == "user" {
		if db.DBVerifyUser(cs[2]) {
			configUser(cs[2])
		} else {
			println("User does not exist.")
		}
	} else if cs[1] == "host" {
		if db.DBVerifyHost(cs[2]) {
			configHost(cs[2])
		} else {
			println("Host does not exist.")
		}
	}
}

func configUser(user string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var c string
		print("[\u001B[1;32m" + user + "\u001B[0;0m]> ")
		c, _ = reader.ReadString('\n')
		c = strings.Split(c, "\n")[0]
		//cs := strings.Split(c, " ")

		if c == "back" {
			return
		} else if c == "options" {
			db.DBUserOptions(user)
		}
	}
}

func configHost(host string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var c string
		print("[\u001B[1;32m" + host + "\u001B[0;0m]> ")
		c, _ = reader.ReadString('\n')
		c = strings.Split(c, "\n")[0]
		//cs := strings.Split(c, " ")

		if c == "back" {
			return
		} else if c == "options" {

		}
	}
}
