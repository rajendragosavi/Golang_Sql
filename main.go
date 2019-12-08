package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("SQL server application")
	db, err := sql.Open("mysql", "DEMOUSER:DEMOPASSWORD@tcp(127.0.0.1:3306)/Demodb")
	//db is an handler for the database
	if err != nil {
		panic(err.Error())
	} 
	//open connection for DB
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error Connecting Database")
		panic(err.Error())
	}

	// Select the database

	_, err = db.Exec("USE Demodb")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Selected the Database Successfully!")
	}

	// To Create a new Table

	_, err = db.Exec("CREATE TABLE Voters(VoterID   INT  NOT NULL, LastName VARCHAR (255)  NOT NULL,FirstName VARCHAR(255)  NOT NULL,State  VARCHAR (255) ,City   VARCHAR (255),  PRIMARY KEY (VoterID) );")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("new table crated successully")
	}

	_, err = db.Exec("INSERT `Voters`(`VoterID`,`LastName`,`FirstName`,`State`,`City`)VALUES('12345','Virat','Kohli','Delhi','Kotla');")
	//Retrive the values from  the table

	rows, err := db.Query("SELECT *FROM Voters")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("")

	column, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	//make a slice to store the column
	values := make([]sql.RawBytes, len(column))
	// rows.scan funcion needs []interfaces{} as input
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	//fetch rows now..

	for rows.Next() {
		//get data now
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		// Now will print columns data
		var data string
		for i, col := range values {
			if col == nil {
				data = "NULL"
			} else {
				data = string(col)
			}
			fmt.Println(column[i], ":", data)

		}
	}
	fmt.Println("---------------------------------------")

	//DELETE the Table
	_, err = db.Exec("DROP TABLE Voters;")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Table deleted successfully!")
	}

	fmt.Println(" Code Passed!")

	defer db.Close()
}
