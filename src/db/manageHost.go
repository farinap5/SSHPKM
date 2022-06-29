package db

import (
	"fmt"
	"github.com/cheynewallace/tabby"
	"log"
)

func DBCreateHost(hostname string) {
	var un string
	row := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", hostname)
	row.Scan(&un)
	if un != "" {
		println("[\u001B[1;31m!\u001B[0;0m] User " + hostname + " already exists.")
	} else {
		sttm, err := DBConn.Prepare(`
		INSERT INTO Host (hostname, name, description, CreateDate,UseToken,Token) VALUES (
		?,?,?,datetime('now','localtime'),?,?) 
		 `)
		if err != nil {
			log.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(hostname, "NULL", "NULL", "false", "NULL")
		if err != nil {
			log.Println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		println("[\u001B[1;32m+\u001B[0;0m] Host " + hostname + " created.")
	}
}

func DBListHost() {
	row, err := DBConn.Query("SELECT hid, Hostname, Name, CreateDate,UseToken FROM Host LIMIT 50")
	if err != nil {
		println(err.Error())
		return
	}

	t := tabby.New()
	t.AddHeader("ID", "Hostname", "Name", "Creation Date", "Authentication")
	for row.Next() {
		var hname, name, cdate, auth string
		var id int
		row.Scan(&id, &hname, &name, &cdate, &auth)
		t.AddLine(id, hname, name, cdate, auth)
	}

	print("\n")
	t.Print()
	print("\n")
}

func DBVerifyHost(name string) bool {
	var un string
	row := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", name)
	row.Scan(&un)
	//println(un)
	if un != "" {
		return true
	} else {
		return false
	}
}

func DBGiveAccess(Host string, User string) {
	va, _ := DBVerifyAccess(User, Host)
	if va {
		fmt.Println("[\u001B[1;31m!\u001B[0;0m] Permission already given.")
		return
	}

	var h, u string
	hrow := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", Host)
	hrow.Scan(&h)
	if h == "" {
		println("[\u001B[1;31m!\u001B[0;0m] Host " + Host + " does not exist.")
		return
	}

	urow := DBConn.QueryRow("SELECT uid FROM User WHERE Username == ?", User)
	urow.Scan(&u)
	if u == "" {
		println("[\u001B[1;31m!\u001B[0;0m] User " + User + " does not exist.")
		return
	}

	sttm, err := DBConn.Prepare(`
	INSERT INTO Access (uid, hid, LastUseDate) VALUES (?,?,datetime('now','localtime'));
	`)
	if err != nil {
		println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return
	}
	_, err = sttm.Exec(u, h)
	if err != nil {
		println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return
	}
	println("[\u001B[1;32m+\u001B[0;0m] Access in " + Host + " given to " + User)
}

func DBListAccess(Host string) {
	var h string
	hrow := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", Host)
	hrow.Scan(&h)
	if h == "" {
		println("[\u001B[1;31m!\u001B[0;0m] Host " + Host + " does not exist.")
		return
	}

	arow, err := DBConn.Query(`
	select Username, A.LastUseDate
		FROM Access A INNER JOIN User u
		ON A.uid = u.uid
		INNER JOIN Host H
		ON A.hid = H.hid
		WHERE H.Hostname = ?;`, Host)
	if err != nil {
		println(err.Error())
		return
	}

	t := tabby.New()
	t.AddHeader("Username", "Last Interaction")
	for arow.Next() {
		var uname, linteract string
		arow.Scan(&uname, &linteract)
		t.AddLine(uname, linteract)
	}
	print("\n")
	t.Print()
	print("\n")
}

func DBVerifyAccess(User string, Host string) (bool, int) {
	arow := DBConn.QueryRow(`
	SELECT a.aid
	FROM Access A INNER JOIN User u
	ON A.uid = u.uid
	INNER JOIN Host H
	ON A.hid = H.hid
	WHERE H.Hostname = ? AND u.Username = ?;
	`, Host, User)
	var a int = 0
	arow.Scan(&a)
	if a == 0 {
		return false, 0
	} else {
		return true, a
	}
}

func DBHostOptions(name string) {
	row := DBConn.QueryRow("SELECT hid, Hostname, Name, Description, CreateDate, UseToken,Token FROM Host WHERE Hostname == ?", name)
	var Hostname, Name, Description, CreateDate, ut, token string
	var hid int
	row.Scan(&hid, &Hostname, &Name, &Description, &CreateDate, &ut, &token)

	t := tabby.New()
	t.AddHeader("Option", "Value")
	t.AddLine("ID", hid)
	t.AddLine("Hostname", Hostname)
	t.AddLine("Creation Date", CreateDate)
	t.AddLine("Name", Name)
	t.AddLine("Description", Description)
	t.AddLine("Need Auth", ut)
	if DBHostNeedAuth(name) {
		t.AddLine("Token", token)
	}
	print("\n")
	t.Print()
	print("\n")
}

func DBSetUpHostVar(v int, value string, host string) {
	switch v {
	case 1: // UPDATE Name
		sttm, err := DBConn.Prepare("UPDATE Host SET Name=? WHERE Hostname=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, host)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		break
	case 2: // UPDATE Desc
		sttm, err := DBConn.Prepare("UPDATE Host SET Description=? WHERE Hostname=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, host)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		break
	case 3: // UPDATE UseAuth
		if value == "true" || value == "false" {

		} else {
			fmt.Println("[\u001B[1;31m!\u001B[0;0m] set auth true or false only.")
			return
		}

		sttm, err := DBConn.Prepare("UPDATE Host SET UseToken=? WHERE Hostname=?;")
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		_, err = sttm.Exec(value, host)
		if err != nil {
			println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
			return
		}
		fmt.Println("[\u001B[1;32m+\u001B[0;0m] auth <- " + value)
		break
	}
}

func DBHostNeedAuth(host string) bool {
	row := DBConn.QueryRow("SELECT UseToken FROM Host WHERE Hostname=?;", host)
	var na string
	row.Scan(&na)
	if na == "false" {
		return false
	} else {
		return true
	}
}

func DBRenewToken(host string) {
	token := DBGenToken(16)
	sttm, err := DBConn.Prepare("UPDATE Host SET Token=? WHERE Hostname=?;")
	if err != nil {
		println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return
	}
	_, err = sttm.Exec(token, host)
	if err != nil {
		println("[\u001B[1;31m!\u001B[0;0m] " + err.Error())
		return
	}
	fmt.Println("[\u001B[1;32m+\u001B[0;0m] Token renewd <- " + token)
}

func DBAuth(token string, host string) bool {
	row := DBConn.QueryRow("SELECT Token FROM Host WHERE Hostname=?;", host)
	var t string
	row.Scan(&t)
	if t == token {
		return true
	} else {
		return false
	}
}

func DBrevoke(User string, Host string) {
	va, aid := DBVerifyAccess(User, Host)
	if !va {
		fmt.Println("[\u001B[1;31m!\u001B[0;0m] No access for user " + User)
		return
	}

	stt, _ := DBConn.Prepare("DELETE FROM Access WHERE aid = ?;")
	stt.Exec(aid)
	fmt.Println("[\u001B[1;32m+\u001B[0;0m] Access revoked for user " + User)
}
