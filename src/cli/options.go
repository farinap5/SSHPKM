package cli

import (
	"fmt"
	"github.com/cheynewallace/tabby"
)

func Banner() {
	fmt.Println("### ###")
}

func Help() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("help", "Help Menu")
	t.AddLine("exit", "Exit.")
	t.AddLine("", "")
	t.AddLine("create", "Create new profile. Type \"create help\" ")
	t.AddLine("list", "List user|host|access <host>")
	t.AddLine("config", "Set variables. Type \"config help\"")
	t.AddLine("access", "Give access for a user to use a host.")
	print("\n")
	t.Print()
	print("\n")
}

func HelpCreate() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("create user <username>", "Create new user.")
	t.AddLine("create host <hostname>", "Create new host.")
	t.AddLine("", "")
	print("\n")
	t.Print()
	print("\nExample: create user test\n\n")

}

func HelpConfig() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("config user <username>", "Set variables of an user.")
	t.AddLine("config host <hostname>", "Set variables of a host.")
	t.AddLine("config", "Typing just `config` local conf will be opened.")
	print("\n")
	t.Print()
	print("\nExample: config user test\n\n")
}

func HelpAccess() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("access <host> <user>", "Give access for a user to use a host.")
	t.AddLine("", "")
	print("\n")
	t.Print()
	print("\nExample: access host1 test\n\n")
}

func OptionsHelp() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("options", "List options and variables")
	t.AddLine("set <option> <value>", "Setup option. e.g. set port 4444")
	t.AddLine("back", "Back to init.")
	print("\n")
	t.Print()
	print("\n")
}
