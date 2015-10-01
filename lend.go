package main

import "os"
import "fmt"
import _ "reflect"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"


var db *sql.DB

func main() {
	args := os.Args[1:]

	initDb()
	defer db.Close()

	if (len(args ) == 0) {
		interactiveMode()
	} else {
		parseArguments(args)
	}
}


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func initDb(){
	var err error

	filename := "./database.db"
	db, err = sql.Open("sqlite3", filename)
	checkErr(err)

	os.Remove(filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// create tables
  		_, err = db.Exec("CREATE TABLE items (id INTEGER PRIMARY KEY, name VARCHAR(255) NOT NULL)")
  		checkErr(err)
  		_, err = db.Exec("CREATE TABLE humans (id INTEGER PRIMARY KEY, name VARCHAR(255) NOT NULL)")
  		checkErr(err)

  		// dummy content
  		createItem("Firefly")
  		createItem("Star wars")
  		createItem("Star wars")
	} 
}


func parseArguments(args []string) {
	switch {
		case len(args) == 0 :
			fmt.Println("Supply at lease one argument")	
		case args[0] == "cat":
			category(args[1:])
		case args[0] == "human":
			human(args[1:])
		case args[0] == "item":
			item(args[1:])
		case args[0] == "switch":
			switchMode(args[1:])
		case args[0] == "lent":
			lent()
		// todo: make case for
		// - lend [item] to [human]
		// - lend [human] returned [item]
	}
}


func createItem(name string) {
	if len(name) > 0 {
		if itemExists(name) {
			fmt.Printf("%1s already exists\n", name)
		} else {
			_, err := db.Exec("INSERT INTO items(name) VALUES(?)", name)
			checkErr(err)
		}	
	}
}


func itemExists(name string) bool {
	var id int
	err := db.QueryRow("SELECT id FROM items WHERE name = ?", name).Scan(&id)

	if err == nil {
        return true
	} else {
        return false
    }
}


func switchMode(args []string) {
	fmt.Println("switch between local and online")	
}


func interactiveMode() {
	fmt.Println("start interactive mode")	
}


func category(args []string) {
	fmt.Println("do something with a category")	
}


func human(args []string) {
	fmt.Println("do something with a human")	
}


func item(args []string) {
	fmt.Println("do something with an item")	
}


func lent() {
	fmt.Println("show list of items that are lent")	
}