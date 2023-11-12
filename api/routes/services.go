package routes

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ngn13/website/api/util"
)

func ServicesDb(db *sql.DB) {
  _, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS services( 
      name    TEXT NOT NULL UNIQUE,
      desc    TEXT NOT NULL,
      url     TEXT NOT NULL
    );
  `)

  if err != nil {
    log.Fatal("Error creating table: "+err.Error())
  }
}

func GetServices(c *fiber.Ctx) error {
  var services []Service = []Service{}

  rows, err := DB.Query("SELECT * FROM services")
  if util.ErrorCheck(err, c) {
    return util.ErrServer(c)
  } 

  for rows.Next() {
    var service Service 
    err := rows.Scan(&service.Name, &service.Desc, &service.Url)
    if err != nil {
      log.Println("Error scaning services row: "+err.Error())
      continue
    }
    services = append(services, service)
  }
  rows.Close()

  return c.JSON(fiber.Map {
    "error": "",
    "result": services, 
  })
}
