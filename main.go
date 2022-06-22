package main

import (
	"github.com/farinap5/SSHPKM/src/cli"
	"github.com/farinap5/SSHPKM/src/db"
)

func main() {
	db.DBFileConfig()
	cli.Start()
}
