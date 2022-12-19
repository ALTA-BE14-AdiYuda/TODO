package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func readConfig() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Error", err.Error())
		return nil
	}
	res := Config{}
	res.DBUser = os.Getenv("DBUSER")
	res.DBPass = os.Getenv("DBPASS")
	res.DBHost = os.Getenv("DBHOST")
	readData := os.Getenv("DBPORT")
	res.DBPort, err = strconv.Atoi(readData)
	if err != nil {
		fmt.Println("Error saat convert", err.Error())
		return nil
	}
	res.DBName = os.Getenv("DBNAME")
	return &res
}

func connectSQL(c Config) *sql.DB {
	// format source username:password@tcp(host:port)/databaseName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := sql.Open("mysql", dsn)
	// db, err := sql.Open("mysql", "Amd A6 5400K:@tcp(localhost:3306/db_tugas")
	if err != nil {
		fmt.Println("terjadi error", err.Error())
	}

	return db
}

type User struct {
	ID       int
	Nama     string
	Password string
}

func main() {
	cfg := readConfig()
	conn := connectSQL(*cfg)
	fmt.Println(conn)

	resultRows, err := conn.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error bos : ", err.Error())
	}

	arrUser := []User{}

	for resultRows.Next() {
		tmp := User{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Password)
		arrUser = append(arrUser, tmp)
	}

	// Exec, Prepare, Query, QueryRow
	// conn.Exec()
	// conn.Prepare()
	// conn.Query()
	// conn.QueryRow()

	fmt.Println(arrUser)
}
