package router

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gralka/simple-fiber-app/common"
	"github.com/gralka/simple-fiber-app/models"
)

func AddBookGroup(app *fiber.App) {
  bookGroup := app.Group("/books")

  bookGroup.Get("/", getBooks)
  bookGroup.Get("/:id", getBook)
  bookGroup.Post("/", createBook)
  bookGroup.Put("/", updateBook)
  bookGroup.Delete("/:id", deleteBook)
}

func getBooks(c *fiber.Ctx) error {
  collection := common.GetDbCollection("books")

  books := make([]models.Book, 0)
  cursor, err := collection.Find(c.Context(), bson.M{})
  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "error": err.Error(),
    })
  }

  for cursor.Next(c.Context()) {
    book := models.Book{}
    err = cursor.Decode(&book)

    if err != nil {
      return c.Status(500).JSON(fiber.Map{
        "error": err.Error(),
      })
    }

    books = append(books, book)
  }

  return c.Status(200).JSON(fiber.Map{ "data": books })
}

func getBook(c *fiber.Ctx) error {
  collection := common.GetDbCollection("books")
  bookId := c.Params("id")

  if bookId == "" {
    return c.Status(400).JSON(fiber.Map{
      "error": "Missing book id",
    })
  }

  objectId, err := primitive.ObjectIDFromHex(bookId)

  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "error": "Invalid book id",
    })
  }

  book := models.Book{}

  err = collection.FindOne(c.Context(), bson.M{ "_id": objectId }).Decode(&book)

  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "error": err.Error(),
    })
  }

  return c.Status(200).JSON(fiber.Map{ "data": book })
}

type createDTO struct {
  Title  string `json:"title" bson:"title"`
  Author string `json:"author" bson:"author"`
  Year   string `json:"year" bson"year"`
}

func createBook(c *fiber.Ctx) error {
  b := new(createDTO)

  if err := c.BodyParser(b); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "error": "Invalid body",
    })
  }

  collection := common.GetDbCollection("books")
  result, err := collection.InsertOne(c.Context(), b)

  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "error": "Failed to create book",
      "message": err.Error(),
    })
  }

  return c.Status(201).JSON(fiber.Map{
    "result": result,
  })
}

type updateDTO struct {
  Title  string `json:"title,omitempty" bson:"title,omitempty"`
  Author string `json:"author,omitempty" bson:"title,omitempty"`
  Year   int    `json:"year,omitempty" bson:"year,omitempty"`
}

func updateBook(c *fiber.Ctx) error {
  b := new(updateDTO)

  if err := c.BodyParser(b); err != nil {
    return c.Status(400).JSON(fiber.Map{
      "error": "Invalid body",
    })
  }

  id := c.Params("id")

  if id == "" {
    return c.Status(400).JSON(fiber.Map{
      "error": "Missing book id",
    })
  }

  objectId, err := primitive.ObjectIDFromHex(id)

  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "error": "Invalid book id",
    })
  }

  collection := common.GetDbCollection("books")
  result, err := collection.UpdateOne(
    c.Context(),
    bson.M{ "_id": objectId },
    bson.M{ "$set": b },
  )

  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "error": "Failed to update book",
      "message": err.Error(),
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "result": result,
  })
}

func deleteBook(c *fiber.Ctx) error {
  id := c.Params("id")

  if id == "" {
    return c.Status(400).JSON(fiber.Map{
      "error": "Missing book id",
    })
  }

  objectId, err := primitive.ObjectIDFromHex(id)

  if err != nil {
    return c.Status(400).JSON(fiber.Map{
      "error": "Invalid book id",
    })
  }

  collection := common.GetDbCollection("books")

  result, err := collection.DeleteOne(c.Context(), bson.M{ "_id": objectId })

  if err != nil {
    return c.Status(500).JSON(fiber.Map{
      "error": "Failed to delete book",
      "message": err.Error(),
    })
  }

  return c.Status(200).JSON(fiber.Map{
    "result": result,
  })
}
