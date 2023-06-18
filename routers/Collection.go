package routers

import (
	"core/database"
	"core/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Collection struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Folders   []Folder   `json:"folders,omitempty"`
	Requests  []Request  `json:"requests,omitempty"`
	Responses []Response `json:"repsonses,omitempty"`
}
type Folder struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Requests     []Request `json:"requests,omitempty"`
	CollectionId int       `json:"collection_id"`
}
type Request struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Data         datatypes.JSON `json:"data"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	FolderId     *int           `json:"folder_id"`
	CollectionId int            `json:"collection_id"`
	Responses    []Response     `json:"repsonses"`
}
type Response struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	Data         datatypes.JSON `json:"data"`
	RequestId    int            `json:"request_id"`
	CollectionId int            `json:"collection_id"`
}

func GetCollection(c *fiber.Ctx) error {
	collections := []models.Collection{}
	db := database.Database.Db.Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return database.Database.Db.Preload("Responses")
		})
	}).Preload("Request", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&collections)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	return c.JSON(collections)
}

func CreateCollect(c *fiber.Ctx) error {
	var collections models.Collection
	if err := c.BodyParser(&collections); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	database.Database.Db.Debug().Create(&collections)
	return c.Status(200).JSON(collections)
}

func DeleteCollection(c *fiber.Ctx) error {
	var collections []Collection
	var folder models.Folder
	var request models.Request
	var response models.Response
	var query Query
	c.QueryParser(&query)
	// if err := c.BodyParser(&collections); err != nil {
	// 	return c.Status(400).JSON(err.Error())

	// }
	database.Database.Db.Where("id = ?", query.ID).First(&collections)
	if len(collections) >= 1 {
		database.Database.Db.Where("collection_id = ?", query.ID).Debug().Delete(response)
		database.Database.Db.Where("collection_id = ?", query.ID).Debug().Delete(request)
		database.Database.Db.Where("collection_id = ?", query.ID).Debug().Delete(folder)
		database.Database.Db.Where("id = ?", query.ID).Debug().Delete(collections)
		return c.Status(200).JSON(fiber.Map{"status": true})
	}
	return c.Status(200).JSON(fiber.Map{"status": false})

}

type Query struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	ExId string `json:"_exporter_id"`
}

func GetCollectionByName(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var collection []models.Collection
	database.Database.Db.Where("name = ? ", query.Name).Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return database.Database.Db.Preload("Response")
		})
	}).Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Response")
	}).Find(&collection)
	return c.Status(200).JSON(collection)
}
func GetCollectionById(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	collection := []models.Collection{}
	//var folder Folder
	db := database.Database.Db.Where("id = ?", query.ID).Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return database.Database.Db.Preload("Responses")
		})
	}).Preload("Request", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&collection)
	if db.Error != nil {
		logrus.Debug("<--------------->", db.Error)
	}
	if len(collection) > 0 {
		return c.Status(200).JSON(collection)
	}
	return c.JSON(fiber.Map{
		"status":  "False",
		"message": "Failed",
		"error":   "Parse data failed",
		"data":    "null",
	})
}
func UpdateCollection(c *fiber.Ctx) error {
	var collections models.Collection
	var col []Collection
	if err := c.BodyParser(&collections); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	var body Collection
	database.Database.Db.Where("id = ?", collections.Id).Find(&col)
	for i := 0; i < len(col); i++ {
		if collections.Id == col[i].Id {
			database.Database.Db.Debug().Where("id = ?", collections.Id).Updates(collections)
			db := database.Database.Db.Preload("Folders", func(db *gorm.DB) *gorm.DB {
				return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
					return database.Database.Db.Preload("Response")
				})
			}).Find(&collections)
			if db.Error != nil {
				logrus.Debug("--------------->", db.Error)
			}
			return c.Status(200).JSON(collections)
		}
	}
	if c.BodyParser(&body.Id) == nil {
		return c.JSON(fiber.Map{
			"status":  "False",
			"message": "Failed",
			"error":   "Parse data failed",
			"data":    "null",
		})
	}
	return c.JSON(fiber.Map{"status": false})

}

type Paramsa struct {
	Id        int
	Name      string
	RequestID int
}

func UpdateCollection1(c *fiber.Ctx) error {
	id := c.Params("id")
	paramsa_id := c.FormValue("id")
	log.Fatal(paramsa_id)
	var collections models.Collection
	var col []Collection
	if err := c.BodyParser(&collections); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	var body Collection
	database.Database.Db.Where("id = ?", id).Find(&col)
	for i := 0; i < len(col); i++ {
		if collections.Id == col[i].Id {
			database.Database.Db.Debug().Where("id = ?", collections.Id).Updates(collections)
			db := database.Database.Db.Preload("Folders", func(db *gorm.DB) *gorm.DB {
				return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
					return database.Database.Db.Preload("Response")
				})
			}).Find(&collections)
			if db.Error != nil {
				logrus.Debug("--------------->", db.Error)
			}
			return c.Status(200).JSON(collections)
		}
	}
	if c.BodyParser(&body.Id) == nil {
		return c.JSON(fiber.Map{
			"status":  "False",
			"message": "Failed",
			"error":   "Parse data failed",
			"data":    "null",
		})
	}
	return c.JSON(fiber.Map{"status": false})

}
