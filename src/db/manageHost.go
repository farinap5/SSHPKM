package db

import (
	"github.com/cheynewallace/tabby"
	"log"
)

func DBCreateHost(hostname string) {
	var un string
	row := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", hostname)
	row.Scan(&un)
	if un != "" {
		println("User " + hostname + " already exists.")
	} else {
		sttm, err := DBConn.Prepare(`
		INSERT INTO Host (hostname, name, description, CreateDate) VALUES (
		?,?,?,datetime('now','localtime')) 
		 `)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = sttm.Exec(hostname, "NULL", "NULL")
		if err != nil {
			log.Println(err.Error())
			return
		}
		println("Host " + hostname + " created.")
	}
}

func DBListHost() {
	row, err := DBConn.Query("SELECT hid, Hostname, Name, CreateDate FROM Host LIMIT 50")
	if err != nil {
		println(err.Error())
		return
	}

	t := tabby.New()
	t.AddHeader("ID", "Hostname", "Name", "Creation Date")
	for row.Next() {
		var hname, name, cdate string
		var id int
		row.Scan(&id, &hname, &name, &cdate)
		t.AddLine(id, hname, name, cdate)
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
	var h, u string
	hrow := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", Host)
	hrow.Scan(&h)
	if h == "" {
		println("Host " + Host + " does not exist.")
		return
	}

	urow := DBConn.QueryRow("SELECT uid FROM User WHERE Username == ?", User)
	urow.Scan(&u)
	if u == "" {
		println("User " + User + " does not exist.")
		return
	}

	sttm, err := DBConn.Prepare(`
	INSERT INTO Access (uid, hid, LastUseDate) VALUES (?,?,datetime('now','localtime'));
	`)
	if err != nil {
		println(err.Error())
		return
	}
	_, err = sttm.Exec(u, h)
	if err != nil {
		println(err.Error())
		return
	}
	println("Access in " + Host + " given to " + User)
}

func DBListAccess(Host string) {
	var h string
	hrow := DBConn.QueryRow("SELECT hid FROM Host WHERE Hostname == ?", Host)
	hrow.Scan(&h)
	if h == "" {
		println("Host " + Host + " does not exist.")
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
	row := DBConn.QueryRow("SELECT hid, Hostname, Name, Description, CreateDate FROM Host WHERE Hostname == ?", name)
	var Hostname, Name, Description, CreateDate string
	var hid int
	row.Scan(&hid, &Hostname, &Name, &Description, &CreateDate)

	t := tabby.New()
	t.AddHeader("Option", "Value")
	t.AddLine("ID", hid)
	t.AddLine("Hostname", Hostname)
	t.AddLine("Creation Date", CreateDate)
	t.AddLine("Name", Name)
	t.AddLine("Description", Description)
	print("\n")
	t.Print()
	print("\n")
}

func DBSetUpHostVar(v int, value string, host string) {
	switch v {
	case 1: // UPDATE Name
		sttm, err := DBConn.Prepare("UPDATE Host SET Name=? WHERE Hostname=?;")
		if err != nil {
			println(err.Error())
			return
		}
		_, err = sttm.Exec(value, host)
		if err != nil {
			println(err.Error())
			return
		}
		break
	case 2: // UPDATE Desc
		sttm, err := DBConn.Prepare("UPDATE Host SET Description=? WHERE Hostname=?;")
		if err != nil {
			println(err.Error())
			return
		}
		_, err = sttm.Exec(value, host)
		if err != nil {
			println(err.Error())
			return
		}
		break
	}
}
