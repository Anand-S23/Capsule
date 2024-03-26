package store

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB(dbUrl string, production bool) *sql.DB {
    db, err := sql.Open("postgres", dbUrl)
    if err != nil {
        log.Fatal("Could not open db :: ", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Could not ping db :: ", err)
    }

    // Note: prod will use db/init.sql file to create the tables if they don't exsits
    if !production {
        err = DropTables(db)
        if err != nil {
            log.Fatal("Could not destory db tables :: ", err)
        }
        log.Println("Dropped all exsiting tables")

        err = CreateTables(db)
        if err != nil {
            log.Fatal("Could not create db tables :: ", err)
            log.Fatal(err)
        }
        log.Println("Recreated all tables")

        err = PopulateTables(db)
        if err != nil {
            log.Fatal("Could not populate db tables :: ", err)
        }
    }

    return db
}

func CreateTables(db *sql.DB) error {
    // Create users table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
            email VARCHAR(255) NOT NULL UNIQUE,
            phone VARCHAR(20) NOT NULL,
            name VARCHAR(255) NOT NULL,
            password VARCHAR(255) NOT NULL
        )
    `)

    if err != nil {
        return err
    }

    return nil
}

func DropTables(db *sql.DB) error {
    // Query to get all table names
    rows, err := db.Query(`
        SELECT table_name
        FROM information_schema.tables
        WHERE table_schema = 'public' AND table_type = 'BASE TABLE';
    `)

    if err != nil {
        return err
    }

    defer rows.Close()

    for rows.Next() {
        var tableName string
        err := rows.Scan(&tableName)
        if err != nil {
            return err
        }

        _, err = db.Exec(fmt.Sprintf("DROP TABLE %s CASCADE;", tableName)) 
        if err != nil {
            return err
        }

        log.Printf("Successfully dropped %s table\n", tableName)
    }

    err = rows.Err()
    return err
}

func PopulateTables(db *sql.DB) error {
    return nil
}

