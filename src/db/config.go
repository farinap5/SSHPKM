package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var DBaseF string = "database.db"

var DBConn *sql.DB

func DBFileConfig() {
	log.Println("Trying to open database file.")
	_, err := os.Open(DBaseF)
	if err != nil {
		log.Println("Database file was not found.")
		log.Println("\t└──Trying to create")
		_, err = os.Create(DBaseF)
		if err != nil {
			log.Println("\t\t└──Fail!")
			panic(err.Error())
		} else {
			log.Println("\t\t└──Success")
		}
	} else {
		log.Println("\t└──Success")
	}
	/*
		Connect
	*/
	log.Println("Trying to connect to the database.")
	tmp, errdb := sql.Open("sqlite3", DBaseF)
	if errdb != nil {
		log.Println("\t\t└──Fail! -> ", err.Error())
	} else {
		log.Println("\t└──Success")
	}
	DBConn = tmp
	DBTesteTables()
}

func DBTesteTables() {
	log.Println("Verifying tables.")

	// Testing table Host
	_, err := DBConn.Query("SELECT * FROM Confing LIMIT 1;")
	if err != nil {
		log.Println("\t└──Fail->", err.Error())
		state, err := DBConn.Prepare(`
		CREATE TABLE IF NOT EXISTS Confing (
		    hid 		INTEGER PRIMARY KEY AUTOINCREMENT,
		    ConfigName	TEXT NOT NULL,
		    Value 		TEXT,
		    Description	TEXT
		);
		`)
		if err != nil {
			log.Println("\t\t└──Fail->", err.Error())
		} else {
			state.Exec()
			log.Println("\t\t└──Success-> Table Confing created.")
		}
	}

	// Testing table User
	_, err = DBConn.Query("SELECT * FROM User LIMIT 1;")
	if err != nil {
		log.Println("\t└──Fail->", err.Error())
		state, err := DBConn.Prepare(`
		CREATE TABLE IF NOT EXISTS User (
		    uid 		INTEGER PRIMARY KEY AUTOINCREMENT,
		    Username 	TEXT NOT NULL,
		    UserID		INTEGER,
		    CreateDate	TEXT NOT NULL,
		    SSHPK		TEXT,
		    Description	TEXT
		);
		`)
		if err != nil {
			log.Println("\t\t└──Fail->", err.Error())
		} else {
			state.Exec()
			log.Println("\t\t└──Success-> Table User created.")
		}
	}

	// Testing table Host
	_, err = DBConn.Query("SELECT * FROM Host LIMIT 1;")
	if err != nil {
		log.Println("\t└──Fail->", err.Error())
		state, err := DBConn.Prepare(`
		CREATE TABLE IF NOT EXISTS Host (
		    hid 		INTEGER PRIMARY KEY AUTOINCREMENT,
		    Hostname	TEXT NOT NULL,
		    Name 		TEXT,
		    Description	TEXT,
		    CreateDate	TEXT NOT NULL
		);
		`)
		if err != nil {
			log.Println("\t\t└──Fail->", err.Error())
		} else {
			state.Exec()
			log.Println("\t\t└──Success-> Table Host created.")
		}
	}

	// Testing table Access
	_, err = DBConn.Query("SELECT * FROM Access LIMIT 1;")
	if err != nil {
		log.Println("\t└──Fail->", err.Error())
		state, err := DBConn.Prepare(`
		CREATE TABLE IF NOT EXISTS Access (
			aid		 		INTEGER PRIMARY KEY AUTOINCREMENT,
			uid             INTEGER,
        	hid             INTEGER,
        	LastUseDate     TEXT,
        	FOREIGN KEY(uid) REFERENCES User (uid),
        	FOREIGN KEY(hid) REFERENCES Host (hid)
		);
		`)
		if err != nil {
			log.Println("\t\t└──Fail->", err.Error())
		} else {
			state.Exec()
			log.Println("\t\t└──Success-> Table Access created.")
		}
	}
}
