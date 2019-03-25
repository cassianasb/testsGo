package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:adm123!@/cursogo")
	if nil != err {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, errs := db.Prepare("update usuarios set nome = ? where id = ?")

	if nil != errs {
		log.Fatal(errs)
	}

	stmt.Exec("Ana", 1)
	stmt.Exec("Paula", 2)

	stmt2, errs2 := db.Prepare("delete from usuarios where id = ?")

	if nil != errs2 {
		log.Fatal(errs2)
	}

	stmt2.Exec(3)
}
