package postgresql

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

type card struct{
    id int
    card_name string
    arcana string
    position string
    yes_no_answer string
}

func GetYesNoAnswer() String {
  connStr := "user=postgres password= dbname= sslmode=disable"
  db, err := sql.Open("postgres", connStr)
  if err != nil {
    panic(err)
    }
  defer db.Close()

  rows, err := db.Query("select * from Products")
    if err != nil {
        panic(err)
    }
  defer rows.Close()
  cards := []card{}

}
