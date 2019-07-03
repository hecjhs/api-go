package api

import (
	"database/sql"
)

func RunFixture() error {
	database, _ := sql.Open("sqlite3", "./data.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS domain (id INTEGER PRIMARY KEY, domain TEXT priority INTEGER, weight INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO domain (priority, weight) VALUES (?, ?)")
	statement, _ = database.Prepare("INSERT INTO domain (priority, weight) VALUES (?, ?)")
	statement, _ = database.Prepare("INSERT INTO domain (priority, weight) VALUES (?, ?)")
	statement.Exec("alpha", 5, 5)
	statement.Exec("omega", 1, 5)
	statement.Exec("beta", 5, 1)
	return nil
}
