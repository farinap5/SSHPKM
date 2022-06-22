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
	row := DBConn.QueryRow("SELECT * FROM Host WHERE Hostname == ?", name)
	row.Scan(&un)
	if un != "" {
		return true
	} else {
		return false
	}
}
