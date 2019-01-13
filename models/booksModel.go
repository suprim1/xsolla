package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Book struct{
	Id int
	Name string
}

func getDbOpen()(*sql.DB, error){
	db, err := sql.Open("mysql", "root:@/golangdb")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func GetBooks()[]Book{
	db, err := getDbOpen()

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	bks := []Book{}
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.Id,&bk.Name)

		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return bks
}

func SetBook(name string)(error){
	db, err := getDbOpen()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("insert into books (name) values (?)", name)
	return err
}

func GetBook(id string)Book{
	db, err := getDbOpen()
	rows:= db.QueryRow("SELECT * FROM books WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	bk := Book{}
	error := rows.Scan(&bk.Id,&bk.Name)
	if error != nil {
		log.Fatal(err)
	}
	return bk
}

func UpdateBook(id string, name string){
	db, err := getDbOpen()
	_, err = db.Exec("update books set name = ? where id = ?",name, id)

	if err != nil {
		log.Println(err)
	}
}

func Del(id string){
	db, err := getDbOpen()
	_, err = db.Exec("delete from books where id = ?", id)
	if err != nil{
		log.Println(err)
	}
}