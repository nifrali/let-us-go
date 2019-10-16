package sqltest

import (
	"database/sql"
	"errors"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb" //because
)

//TestDbConnection returns a string representing the SQL server version if a connection can be made
func TestDbConnection(username, password, hostname, database, instance, port string) (string, error) {
	query := url.Values{}
	query.Add("database", database)
	query.Add("port", port)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     hostname,
		RawQuery: query.Encode(),
	}
	if instance != "" {
		u.Path = instance
	}

	condb, errdb := sql.Open("mssql", u.String())
	if errdb != nil {
		return "", errdb
	}

	defer condb.Close()

	var sqlversion string

	rows, err := condb.Query("select @@version")
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		return sqlversion, nil
	}
	return "", errors.New("No rows returned")
}
