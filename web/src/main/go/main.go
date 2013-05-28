package main

import (
//  _ "github.com/go-sql-driver/mysql"
//  "database/sql"
  "fmt"
  "net/http"
  "handlers"
)

/*
func readDb(w http.ResponseWriter, r *http.Request) {
  db, err := sql.Open("mysql", "root:mtrpls12@unix(/var/run/mysqld/mysqld.sock)/broke?charset=utf8")

  if err != nil {
    panic(err)
  }
  defer db.Close()

  row := db.QueryRow("SELECT `Id`, `UserName`, `EmailAddress`, `Password`, `IsActive`, `DateJoined`, `LastLoggedOn`, `IsAdministrator` FROM `Person` WHERE Id=?;", id)
  var person Person
  row.Scan(&person.Id, &person.UserName, &person.EmailAddress, &person.Password, &person.IsActive, &person.DateJoined, &person.LastLoggedOn, &person.IsAdministrator)
}
*/

func main() {
  port := ":8080"
  http.HandleFunc("/", echo_name.handler)
  fmt.Printf("Server starting on port %s", port)
  http.ListenAndServe(port, nil)
}
