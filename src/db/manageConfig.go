package db

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
