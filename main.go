package main

import (
	"core/database"
	"core/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	fmt.Println("hello")
	database.ConnectDB()
	SetUpRouter(app)
	app.Listen(":8000")
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Range",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,DELETE,PUT",
		ExposeHeaders:    "X-Total-Count, Content-Range",
	}))
}

func SetUpRouter(app *fiber.App) {

	app.Use(cors.New())
	app.Group("/", logger.New())
	//collection
	app.Get("/collection", routers.GetCollection)
	app.Post("/collection", routers.CreateCollect)
	app.Put("/collection/:id", routers.UpdateCollection1)
	app.Put("/collection", routers.UpdateCollection)
	app.Delete("/collection", routers.DeleteCollection)
	app.Get("/collection_id", routers.GetCollectionById)
	//folder
	app.Get("/folder", routers.GetFolder)
	app.Post("/folder", routers.CreateFolder)
	app.Delete("/folder", routers.DeleteFolder)
	app.Put("/folder", routers.UpdateFolder)
	app.Get("/folder_id", routers.GetFolderById)
	app.Get("/folder_id_collection", routers.GetFolderByIdCollection)

	//Request
	app.Get("/request", routers.GetRequest)
	app.Post("/request", routers.CreateRequest)
	app.Delete("/request", routers.DeleteRequest)
	app.Put("/request", routers.UpdateRequest)
	app.Get("/request_id", routers.GetRequestById)

	//Response
	app.Get("/response", routers.GetResponse)
	app.Post("/response", routers.CreateResponse)
	app.Delete("/response", routers.DeleteResponse)
	app.Put("/response", routers.UpdateResponse)
	app.Get("/response_id", routers.GetResponseById)

	//mock_server
	mock_server := app.Group("/mock_server", logger.New())
	mock_server.Get("/:name/:path", routers.Method)
	mock_server.Post("/:name/:path", routers.Method)
	mock_server.Put("/:name/:path", routers.Method)
	mock_server.Patch("/:name/:path", routers.Method)
	mock_server.Delete("/:name/:path", routers.Method)

	mock_server.Get("/export", routers.Export2)
	mock_server.Post("/import", routers.Import)
}
