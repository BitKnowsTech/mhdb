package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"

	_ "github.com/mattn/go-sqlite3"
)

const DATABASEPATH = "./build/database"
const DATABASEFILE = "mhdb.db"
const DATADIR = "./data"

func main() {
	fmt.Println("[INITIALIZING DB]")
	db := InitDB()
	defer db.Close()
	fmt.Println("[CREATING MONSTERS]")
	createMonsters(db)
}

func InitDB() *sql.DB {
	if _, err := os.Stat(DATABASEPATH); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(DATABASEPATH, 0755); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Removing old Database...")
	os.Remove(DATABASEPATH + "/" + DATABASEFILE)

	fmt.Println("Creating new Database...")
	db, err := sql.Open("sqlite3", DATABASEPATH+"/"+DATABASEFILE)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func createMonsters(db *sql.DB) {
	type Monster struct {
		Id      int    `csv:"id"`
		Name    string `csv:"name"`
		IsLarge int    `csv:"is_large"`
	}
	var monsters []Monster

	monsterFile, err := os.Open(fmt.Sprintf("%s/monsters.csv", DATADIR))
	if err != nil {
		log.Fatal(err)
	}
	defer monsterFile.Close()

	if err := gocsv.Unmarshal(monsterFile, &monsters); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Monster Table...")
	createTableSQL := `CREATE TABLE IF NOT EXISTS monsters (
        id INTEGER PRIMART KEY,
        name TEXT NOT NULL,
        is_large BOOL NOT NULL
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Inserting Monsters...")
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO monsters (id, name, is_large) VALUES (?, ?, ?)")
	defer stmt.Close()

	for _, monster := range monsters {
		stmt.Exec(monster.Id, monster.Name, monster.IsLarge)
	}
	tx.Commit()
	fmt.Println("DONE")

}
