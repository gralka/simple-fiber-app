package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverMw "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gralka/simple-fiber-app/common"
	"github.com/gralka/simple-fiber-app/router"
)

func main() {
  err := run()

  if err != nil {
    panic(err)
  }
}

func run() error {
  // Load environment variables
  err := common.LoadEnv()

  if err != nil {
    return err
  }
 
  // Initialize the database connection
  err = common.InitDb() 

  if err != nil {
    return err
  }

  // Defer closing the database connection
  defer common.CloseDb()

  // Create a new fiber application.
  app := fiber.New()

  // Add middleware
  app.Use(logger.New())
  app.Use(recoverMw.New())
  app.Use(cors.New())

  router.AddBookGroup(app)

  var port string

  if port = os.Getenv("PORT"); port == "" {
    port = "8080"
  }

  app.Listen(":" + port)

  return nil
}
