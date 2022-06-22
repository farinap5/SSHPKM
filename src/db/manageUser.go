package db

import (
	"fmt"
	"github.com/cheynewallace/tabby"
	"log"
)

func DBCreateUser(name string) {
	var un string
	row := DBConn.QueryRow("SELECT uid FROM User WHERE Username == ?", name)
	row.Scan(&un)
	if un != "" {
		println("User " + name + " already exists.")
	} else {
		sttm, err := DBConn.Prepare(`
		INSERT INTO User (username, userid, createdate, sshpk, description) VALUES (
		?,?,datetime('now','localtime'),?,?) 
		 `)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = sttm.Exec(name, 0, "NULL", "NULL")
		if err != nil {
			log.Println(err.Error())
			return
		}
		println("User " + name + " created.")
	}
}

func DBListUser() {
	row, err := DBConn.Query("SELECT * FROM User LIMIT 50")
	if err != nil {
		println(err.Error())
		return
	}

	t := tabby.New()
	t.AddHeader("ID", "Username", "UID", "Creation Date", "SSH PK")
	for row.Next() {
		var uname, cdate, sshpk, desc string
		var id, uid int
		row.Scan(&id, &uname, &uid, &cdate, &sshpk, &desc)
		if sshpk == "NULL" {
			sshpk = "false"
		} else {
			sshpk = "true"
		}
		t.AddLine(id, uname, uid, cdate, sshpk)
	}

	print("\n")
	t.Print()
	print("\n")
}

func DBVerifyUser(name string) bool {
	var un string
	row := DBConn.QueryRow("SELECT uid FROM User WHERE Username == ?", name)
	row.Scan(&un)
	if un != "" {
		return true
	} else {
		return false
	}
}

func DBUserOptions(name string) {
	row := DBConn.QueryRow("SELECT * FROM User WHERE Username == ?", name)
	var uname, cdate, sshpk, desc string
	var id, uid int
	row.Scan(&id, &uname, &uid, &cdate, &sshpk, &desc)
	fmt.Printf("\n%d) %s\n\n", id, uname)
	fmt.Printf("Creation Date: %s\n\n", cdate)
	fmt.Printf("SSH Public Key:\n%s\n\n", sshpk)
	fmt.Printf("User Description:\n%s\n\n", desc)
}
