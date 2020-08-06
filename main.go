package main

import (
	"database/sql"
	"fmt"
	"log"

	//im gonna import it but im not gonna reference it thats why i use underscore
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//to know wish driver im working in
	fmt.Println(sql.Drivers())
	//connexion to db
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangportfolio")
	if err != nil {
		log.Fatal("Unable to open connexion to database")
	}
	//close db in the end of operation using defer
	defer db.Close()
	/*//////////////////////select table///////////////////////////*/
	//new query select from table test
	results, err := db.Query("select * from test")
	if err != nil {
		log.Fatal("error when fetching golangportfolio", err)
	}
	//when i get result and i finish whit it i need to close
	defer results.Close()
	//to loop over this result sequentially
	for results.Next() {
		//value that ill receive from db
		var (
			id       int
			lastName string
			name     string
		)
		// i usee scane to get data from db number of argument is the some of columns returned or run time wil receive an error
		//use pointers to store value on it
		err = results.Scan(&id, &name, &lastName)
		if err != nil {
			log.Fatal("unablle to parse row: ", err)
		}
		fmt.Printf("id: %d, lastName: %s, name: %s\n", id, lastName, name)

	}
	/*//////////////////////select by id///////////////////////////*/
	var (
		id             int
		lastName, name string
	)
	// new query it return one row if you receive more than one it whill take the first and throw the rest
	err = db.QueryRow("select * from test where id = 1").Scan(&id, &lastName, &name)
	if err != nil {
		log.Fatal("unab to parse row", err)
	}
	fmt.Printf("id: %d, lastName: %s, name: %s\n", id, lastName, name)
	/*//////////////////////insert///////////////////////////*/
	//slice of struct that got a lot of value
	persons := []struct {
		lastname, name string
	}{
		{"chaima", "elgouail"},
		{"sara", "elgouail"},
		{"taher", "mazizi"},
	}
	//prepare our statement
	insertForm, err := db.Prepare("INSERT INTO test(name,lastName) VALUES(?,?)")
	//close statement when i finish whit it
	defer insertForm.Close()
	if err != nil {
		log.Fatal("Unable to prepare statement", err)
	}
	for _, person := range persons {
		//the number of argument need to be the same as the number of place holder VALUES(?,?)
		_, err = insertForm.Exec(person.name, person.lastname)
		if err != nil {
			log.Fatal("Unable to execute statement ", err)
		} else {
			fmt.Printf("name: %s,lastName: %s", person.name, person.lastname)
			fmt.Println(" has been inserted successfully")
		}

	}
	/*//////////////////////delete///////////////////////////*/
	deletForm, err := db.Prepare("DELETE FROM test WHERE id=? ")

	if err != nil {
		log.Fatal("Unable to prepare deleteform : ", err)
	}
	deletForm.Exec(2)
	deletForm.Close()
	fmt.Printf("row whit id= %d has been deleted successfully", 2)
	fmt.Println()
	/*//////////////////////Update///////////////////////////*/
	updateForm, err := db.Prepare("update test set name=? where id =?")
	defer updateForm.Close()
	if err != nil {
		log.Fatal("Unable to repare updateForm ", err)
	}
	updateForm.Exec("chadiUpdated", 3)
	fmt.Printf("row whit id %d has benn update it successfully", 3)
}
