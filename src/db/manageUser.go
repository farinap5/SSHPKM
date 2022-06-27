package db

import (
	"github.com/cheynewallace/tabby"
	"log"
	"strings"
)

func DBCreateUser(name string) {
	var un string
	row := DBConn.QueryRow("SELECT uid FROM User WHERE Username == ?", name)
	row.Scan(&un)
	if un != "" {
		println("[\u001B[1;31m!\u001B[0;0m] User " + name + " already exists.")
	} else {
		sttm, err := DBConn.Prepare(`
		INSERT INTO User (username, userid, createdate, sshpk, description) VALUES (
		?,?,datetime('now','localtime'),?,?) 
		 `)
		if err != nil {
			log.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(name, 0, "NULL", "NULL")
		if err != nil {
			log.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		println("[\u001B[1;32m+\u001B[0;0m] User " + name + " created.")
	}
}

func DBListUser() {
	row, err := DBConn.Query("SELECT * FROM User LIMIT 50")
	if err != nil {
		println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
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
	row := DBConn.QueryRow("SELECT uid, Username, UserID, CreateDate, SSHPK, Description FROM User WHERE Username == ?", name)
	var uname, cdate, sshpk, desc string
	var id, uid int
	row.Scan(&id, &uname, &uid, &cdate, &sshpk, &desc)

	sshpks := strings.Split(sshpk, " ")

	t := tabby.New()
	t.AddHeader("Option", "Value")
	t.AddLine("ID", id)
	t.AddLine("Username", uname)
	t.AddLine("User ID", uid)
	t.AddLine("Creation Date", cdate)
	if sshpks[0] != "NULL" {
		t.AddLine("Public Key", sshpks[0]+" "+sshpks[2])
	} else {
		t.AddLine("Public Key", sshpks[0])
	}
	t.AddLine("Description", desc)
	print("\n")
	t.Print()
	print("\n")
}

func DBSetuUserVar(v int, value string, name string) {
	switch v {
	case 1: // UPDATE SSHKEY
		sttm, err := DBConn.Prepare("UPDATE User SET SSHPK=? WHERE Username=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, name)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		}
		break

	case 2: // UPDATE Description
		sttm, err := DBConn.Prepare("UPDATE User SET Description=? WHERE Username=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, name)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		break

	case 3: // UPDATE UID
		sttm, err := DBConn.Prepare("UPDATE User SET UserID=? WHERE Username=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, name)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		}
		break
	}
}

func DBGetUserSSHKey(User string) string {
	var key string
	row := DBConn.QueryRow("SELECT SSHPK FROM User WHERE Username == ?", User)
	row.Scan(&key)
	if key == "" {
		return "(null)"
	} else {
		return key
	}
}
