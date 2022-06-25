package db

import (
	"fmt"
	"github.com/cheynewallace/tabby"
)

func DBGetConfig(Conf string) string {
	var c string
	hrow := DBConn.QueryRow("SELECT Value FROM Confing WHERE ConfigName == ?", Conf)
	hrow.Scan(&c)
	if c == "" {
		println("Not conf.")
		return "NULL"
	}
	return c
}

func DBListConfig() {
	crow, err := DBConn.Query("SELECT hid,ConfigName,Value,Description FROM Confing;")
	if err != nil {
		fmt.Println(err.Error())
	}
	t := tabby.New()
	t.AddHeader("ID", "Config Name", "Value", "Description")
	for crow.Next() {
		var hid int
		var cname, value, desc string
		crow.Scan(&hid, &cname, &value, &desc)
		t.AddLine(hid, cname, value, desc)
	}
	print("\n")
	t.Print()
	print("\n")
}

func DBSetConf(option string, value string) bool {
	if option == "Address" || option == "address" || option == "addr" {
		option = "Address"
	} else if option == "Port" || option == "port" || option == "p" {
		option = "Port"
	} else if option == "Log" || option == "log" || option == "logs" {
		option = "Logs"
	} else {
		return false
	}

	sttm, err := DBConn.Prepare("UPDATE Confing SET Value=? WHERE ConfigName=?;")
	if err != nil {
		println(err.Error())
		return false
	}
	_, err = sttm.Exec(value, option)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}
