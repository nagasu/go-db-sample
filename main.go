// 参考サイト https://docs.microsoft.com/ja-jp/azure/postgresql/connect-go

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DB設定値
const (
    HOST     = "localhost"
    DATABASE = "go-sample"
    USER     = "postgres"
    PASSWORD = "postgres"
)

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    create()
}

func create() {
    connectionString := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s sslmode=disable",
        HOST,
        USER,
        PASSWORD,
        DATABASE)

    // Initialize connection object.
    db, err := sql.Open("postgres", connectionString)
    checkError(err)

    err = db.Ping()
    checkError(err)
    fmt.Println("Successfully created connection to database")

    // Drop previous table of same name if one exists.
    _, err = db.Exec("DROP TABLE IF EXISTS inventory;")
    checkError(err)
    fmt.Println("Finished dropping table (if existed)")

    // Create table.
    _, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
    checkError(err)
    fmt.Println("Finished creating table")

    // Insert some data into table.
    sqlStatement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
    _, err = db.Exec(sqlStatement, "banana", 150)
    checkError(err)
    _, err = db.Exec(sqlStatement, "orange", 154)
    checkError(err)
    _, err = db.Exec(sqlStatement, "apple", 100)
    checkError(err)
    fmt.Println("Inserted 3 rows of data")
}
