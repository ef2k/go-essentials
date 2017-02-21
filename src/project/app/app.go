package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	// The database connection that this application should use.
	DB *sql.DB
)

// The entry point to the application, starts up the server on the given port.
func ListenAndServe(port string) error {
	http.HandleFunc("/favicon.ico", handleFavicon)
	http.HandleFunc("/health-check", handleHealthCheck)
	http.HandleFunc("/", handleIndex)
	return http.ListenAndServe(":"+port, nil)
}
