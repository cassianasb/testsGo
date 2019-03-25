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

	tx, _ := db.Begin() //Iniciar uma transação
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(1, "Tiago") //chave duplicada

	if nil != err { //sem o tratamento do erro, os usuarios sem chave duplicada seriam criados
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
