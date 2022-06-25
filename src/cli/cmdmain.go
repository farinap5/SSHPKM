package cli

import (
	"bufio"
	"fmt"
	"github.com/farinap5/SSHPKM/src/db"
	"github.com/farinap5/SSHPKM/src/server"
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
		server.ServerStop()
		os.Exit(0)
	} else if cs[0] == "create" {
		create(cs)
	} else if cs[0] == "list" {
		list(cs)
	} else if cs[0] == "config" {
		config(cs)
	} else if cs[0] == "access" {
		if cs[1] == "help" {
			HelpAccess()
			return
		}
		if len(cs) != 3 {
			println("More arguments needed.")
			return
		}
		db.DBGiveAccess(cs[1], cs[2])
	} else if cs[0] == "listen" || cs[0] == "server" {
		listen(cs)
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
	} else if cs[1] == "access" && len(cs) == 3 {
		db.DBListAccess(cs[2])
	}
}

func config(cs []string) {
	if len(cs) == 1 {
		configLocal()
		return
	}

	if cs[1] == "help" && len(cs) > 1 {
		HelpConfig()
		return
	}

	if len(cs) != 3 {
		fmt.Println("More arguments needed.")
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
		cs := strings.Split(c, " ")

		if c == "back" {
			return
		} else if c == "options" {
			db.DBUserOptions(user)
		} else if cs[0] == "set" && len(cs) == 3 {
			if cs[1] == "pk" {
				db.DBSetuUserVar(1, cs[2], user)
				fmt.Println("Public Key <- " + cs[2])
			} else if cs[1] == "uid" {
				db.DBSetuUserVar(3, cs[2], user)
				fmt.Println("UserID <- " + cs[2])
			} else if cs[1] == "desc" {
				db.DBSetuUserVar(2, cs[2], user)
				fmt.Println("Description <- " + cs[2])
			}
		} else if c == "help" {
			OptionsHelp()
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
		cs := strings.Split(c, " ")

		if c == "back" {
			return
		} else if c == "options" {
			db.DBHostOptions(host)
		} else if c == "help" {
			OptionsHelp()
		} else if cs[0] == "set" && len(cs) == 3 {
			if cs[1] == "name" {
				db.DBSetUpHostVar(1, cs[2], host)
				fmt.Println("Name <- " + cs[2])
			} else if cs[1] == "desc" {
				db.DBSetUpHostVar(2, cs[2], host)
				fmt.Println("Descriptions <- " + cs[2])
			}
		}
	}
}

func configLocal() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var c string
		print("[\u001B[1;32mLocal\u001B[0;0m]> ")
		c, _ = reader.ReadString('\n')
		c = strings.Split(c, "\n")[0]
		cs := strings.Split(c, " ")

		if c == "back" {
			break
		} else if c == "options" {
			db.DBListConfig()
		} else if c == "help" {
			OptionsHelp()
		} else if cs[0] == "set" && len(cs) == 3 {
			if db.DBSetConf(cs[1], cs[2]) {
				fmt.Println("OK " + cs[1] + " <- " + cs[2])
			} else {
				fmt.Println("Error: " + cs[1] + "<-" + cs[2])
			}
		} else if c == "back" {
			return
		}
	}
}

func listen(cs []string) {
	if len(cs) != 2 {
		println("More arguments needed.")
		return
	}

	if cs[1] == "start" {
		server.ServerListen()
		CMDStatus = "\033[1;32m*\033[0;0m"
	} else if cs[1] == "stop" || cs[1] == "shutdown" {
		if server.ServerStop() {
			CMDStatus = "\033[1;31m*\033[0;0m"
		}
	} else if cs[1] == "restart" {
		server.ServerRestart()
	}
}
