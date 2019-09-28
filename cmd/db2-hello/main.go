package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	//con := "HOSTNAME=host;DATABASE=name;PORT=number;UID=username;PWD=password"
	con := "HOSTNAME=192.168.1.68;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=mypwd"
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {

		fmt.Println(err)
	}
	defer db.Close()

	if err := create(db); err != nil {
		log.Fatal(err)
	}

	if err := insert(db); err != nil {
		log.Fatal(err)
	}

	if err := display(db); err != nil {
		log.Fatal(err)
	}
}

func create(db *sql.DB) error {
	_, err := db.Exec("DROP table SAMPLE")
	if err != nil {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec("create table SAMPLE(ID varchar(20),NAME varchar(20),LOCATION varchar(20),POSITION varchar(20))")
		if err != nil {
			return err
		}
	}
	fmt.Println("TABLE CREATED")
	return nil
}

func insert(db *sql.DB) error {
	st, err := db.Prepare("Insert into SAMPLE(ID,NAME,LOCATION,POSITION) values('3242','mike','hyd','manager')")
	if err != nil {
		return err
	}
	st.Query()
	return nil
}

func display(db *sql.DB) error {
	st, err := db.Prepare("select * from SAMPLE")
	if err != nil {
		return err
	}
	err = execquery(st)
	if err != nil {
		return err
	}
	return nil
}

func execquery(st *sql.Stmt) error {
	rows, err := st.Query()
	if err != nil {
		return err
	}
	cols, _ := rows.Columns()
	fmt.Printf("%s    %s   %s    %s\n", cols[0], cols[1], cols[2], cols[3])
	fmt.Println("-------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var t, x, m, n string
		err = rows.Scan(&t, &x, &m, &n)
		if err != nil {
			return err
		}
		fmt.Printf("%v  %v   %v         %v\n", t, x, m, n)
	}
	return nil
}
