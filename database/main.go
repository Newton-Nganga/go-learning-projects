package main

import (
	"database/sql"
	"fmt"
	"log"
	//_ "github.com/go-sql-driver/mysql"
)

type User struct{
	ID int
	Name string
	Age int
}

func main(){
  db,err := sql.Open("mysql","username:password@tcp(localhost:3306/database_name)")
  if err != nil{
	log.Fatal(err)
  }
  defer db.Close()


  //Query users
  rows,err := db.Query("SELECT id,name,age FROM users")
  if err != nil{
	log.Fatal(err)
  }
  
  var users []User

  for rows.Next(){
	var user User
	err := rows.Scan(&user.ID,&user.Name,&user.Age)
	if err != nil{
		log.Fatal(err)
	}
	users = append(users,user)
  }
  if err := rows.Err(); err != nil {
	log.Fatal(err)
  }

  for _ , user := range users {
	fmt.Println(user.ID,user.Name,user.Age)
  }



}

