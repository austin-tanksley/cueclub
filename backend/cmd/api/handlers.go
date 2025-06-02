package main

import (
	"cueclub/data"
	"database/sql"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func getMix(c *fiber.Ctx) error {
	db, err := sql.Open("sqlite3", "./backend/data/cueclub.db")
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM mixes").Scan(&count)
	if err != nil {
		return fiber.DefaultErrorHandler(c, err)
	}
	randomID := rand.Intn(count) + 1

	row := db.QueryRow("SELECT * FROM mixes WHERE id = ?", randomID)
	var m data.Mix
	err = row.Scan(&m.ID, &m.YTID, &m.URL, &m.Title, &m.Creator, &m.CreatorURL, &m.Version)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(m, "application/json")
}
