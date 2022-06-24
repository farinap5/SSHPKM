package db

type User struct {
	uid         int
	username    string
	userid      int
	createdate  string
	sshpk       string
	description string
}

type Host struct {
	hid        int
	hostname   string
	name       string
	createdate string
}
