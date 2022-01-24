package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type ticket struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func getTickets(c *gin.Context) {
	var tickets = []ticket{}
	var rows *sql.Rows
	rows, err = db.Query("select * from tickets")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// See if you can get number of records already.
	i := 0
	for rows.Next() {
		tickets = append(tickets, ticket{ID: "0", Title: " "})
		err := rows.Scan(&tickets[i].ID, &tickets[i].Title)
		if err != nil {
			panic(err)
		}
		i++
	}

	c.IndentedJSON(http.StatusOK, tickets)
}

func getTicketID(c *gin.Context) {
	var requestedTicket ticket
	var row *sql.Rows

	id := c.Param("id")
	fmt.Println(id)
	row, err = db.Query("select id, title from tickets where id = ?", id)
	if err != nil {
		return
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&requestedTicket.ID, &requestedTicket.Title)
		if err != nil {
			panic(err)
		}
	}
	c.IndentedJSON(http.StatusOK, requestedTicket)
}

func createTicket(c *gin.Context) {
	var newTicket ticket

	if err := c.BindJSON(&newTicket); err != nil {
		return
	}

	storeTicket(newTicket.ID, newTicket.Title)

}

func create(name string) {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create database if not exists " + name)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use " + name)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists " + name + " ( id varchar(50), title varchar(50))")
	if err != nil {
		panic(err)
	}

}

func storeTicket(id string, title string) {
	sql := fmt.Sprintf(`insert into tickets (id, title) values (%s, '%s')`, id, title)

	fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func main() {
	create("tickets")

	router := gin.Default()

	router.GET("/tickets", getTickets)
	router.GET("/ticket/:id", getTicketID)
	router.POST("/ticket", createTicket)

	router.Run("localhost:8080")
}
