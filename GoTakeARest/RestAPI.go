package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Animal struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Group string `json:"Group"`
	Desc  string `json:"Desc"`
}

func main() {

	// echo instance
	e := echo.New()

	// Middleware
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Animal REST API is running...")
	})

	// Routes
	e.GET("/animals", GetAllAnimals)

	// server address
	e.Logger.Fatal(e.Start(":4200"))
}

func GetAllAnimals(c echo.Context) error {
	// open database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// query
	rows, err := db.Query("SELECT ID, Name, Group, Desc FROM Animal WHERE Name IS NOT NULL AND ID IS NOT NULL AND Group IS NOT NULL")
	if err != nil {
		return err
	}
	defer rows.Close()

	// get data
	animals := []Animal{}
	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Group, &animal.Desc)
		if err != nil {
			return err
		}
		animals = append(animals, animal)
	}

	return c.JSON(http.StatusOK, animals)
}
