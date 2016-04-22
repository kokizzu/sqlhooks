/*
Package sqlhooks Attach hooks to any database/sql driver.

The purpose of sqlhooks is to provide anway to instrument your sql statements,
making really easy to log queries or measure execution time without modifying your actual code.

Example:
	package main

	import (
		"database/sql"
		"log"
		"time"

		"github.com/gchaincl/sqlhooks"
		_ "github.com/mattn/go-sqlite3"
	)


	func main() {
		hooks := sqlhooks.Hooks{
			Exec: func(query string, args ...interface{}) func() {
				log.Printf("[exec] %s, args: %v", query, args)
				return nil
			},
			Query: func(query string, args ...interface{}) func() {
				t := time.Now()
				id := t.Nanosecond()
				log.Printf("[query#%d] %s, args: %v", id, query, args)

				// This will be executed when Query statements has completed
				return func() {
					log.Printf("[query#%d] took: %s\n", id, time.Since(t))
				}
			},
		}

		// Register the driver
		// "sqlite-hooked" is the attached driver, and "sqlite3" is where we're attaching to
		sqlhooks.Register("sqlite-hooked", sqlhooks.NewDriver("sqlite3", &hooks))

		// Connect to attached driver
		db, _ := sql.Open("sqlite-hooked", ":memory:")

		// Do you're stuff
		db.Exec("CREATE TABLE t (id INTEGER, text VARCHAR(16))")
		db.Exec("INSERT into t (text) VALUES(?), (?))", "foo", "bar")
		db.Query("SELECT id, text FROM t")
	}

*/
package sqlhooks