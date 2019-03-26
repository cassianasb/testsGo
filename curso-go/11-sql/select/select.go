package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:adm123!@/cursogo")
	if nil != err {
		log.Fatal(err)
	}

	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id < ?", 2)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome) //Mapeando o resultado da consulta.
		fmt.Println(u)
	}
}
