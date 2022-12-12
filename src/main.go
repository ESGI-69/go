package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func ping(context *gin.Context) {
	fmt.Println("GET /")
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// The global database connection
var Database *sql.DB

func main() {
	// Connect to the database
	var databaseError error
	Database, databaseError = sql.Open(
		"mysql",
		os.Getenv("DB_USER")+
			":"+
			os.Getenv("DB_PASSWORD")+
			"@tcp("+
			os.Getenv("DB_HOST")+
			":"+
			os.Getenv("DB_PORT")+
			")/"+
			os.Getenv("DB_NAME"),
	)
	if databaseError != nil {
		fmt.Println(databaseError)
		fmt.Println("❌ Database connection failed")
		panic(databaseError.Error())
	} else if Database.Ping() != nil {
		fmt.Println(Database.Ping())
		fmt.Println("❌ Database connection failed")
		panic(databaseError.Error())
	} else {
		fmt.Println("✅ Database connection established")
	}
	// Close the database connection when the program ends
	defer Database.Close()

	// Create the gin engine
	engine := gin.Default()
	engine.GET("/ping", ping)
	engine.Run()
}
