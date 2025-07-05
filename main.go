package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id         uuid.UUID
	Created_at time.Time
	Updated_at time.Time
	Email      string
}

func main() {

	db, err := sql.Open("sqlite3", "./weigh_in.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// dbqueries := database.New(db)

	// Initialize db tables
	// TODO: write function for initial schema migration

	_, err = db.Exec("create table if not exists users (id uuid primary key, created_at timestamp, updated_at timestamp, email text not null unique);")
	if err != nil {
		log.Fatal(err)
	}

	// new_user, err := dbqueries.CreateUser(context.Background(), database.CreateUserParams{
	// 	ID:        uuid.New(),
	// 	Email:     "mkowalsky87@gmail.com",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Enjoy Not Sucking",
		})
	})

	// Kamal healthcheck route
	router.GET("/up/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run(":80")
}
