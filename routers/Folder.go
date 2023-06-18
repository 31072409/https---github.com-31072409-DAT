package routers

import (
	"core/database"
	"core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Folder1 struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Requests     []Request `json:"requests,omitempty"`
	CollectionId int       `json:"collection_id"`
}
type Request1 struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Data         datatypes.JSON `json:"data"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	CollectionId int            `json:"collection_id"`
	FolderId     int            `json:"folder_id"`
	Responses    []Response     `json:"repsonses"`
}
type Response1 struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	Data         datatypes.JSON `json:"data"`
	RequestId    int            `json:"request_id"`
	CollectionId int            `json:"collection_id"`
}

func GetFolder(c *fiber.Ctx) error {
	folder := []models.Folder{}
	db := database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&folder)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	return c.JSON(folder)
}

func CreateFolder(c *fiber.Ctx) error {
	var folder models.Folder
	if err := c.BodyParser(&folder); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Debug().Create(&folder)
	return c.Status(200).JSON(folder)
}

func DeleteFolder(c *fiber.Ctx) error {
	var folder []Folder
	var request []Request
	var response models.Response
	var query Query
	c.QueryParser(&query)
	database.Database.Db.Where("folder_id = ?", query.ID).Find(&request)
	database.Database.Db.Where("id = ?", query.ID).Find(&folder)

	if len(folder) >= 1 {
		for i := 0; i < len(request); i++ {
			database.Database.Db.Where("request_id = ?", request[i].Id).Debug().Delete(response)
		}
		database.Database.Db.Where("folder_id = ?", query.ID).Debug().Delete(request)
		database.Database.Db.Where("id = ?", query.ID).Debug().Delete(folder)

		return c.Status(200).JSON(fiber.Map{"status": true})
	}
	return c.Status(200).JSON(fiber.Map{"status": false})
}
func GetFolderByName(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var folder []models.Folder
	database.Database.Db.Where("name = ? ", query.Name).Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return database.Database.Db.Preload("Response")
		})
	}).Find(&folder)
	return c.Status(200).JSON(folder)
}
func GetFolderById(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var folder []models.Folder
	db := database.Database.Db.Where("id = ?", query.ID).Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&folder)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	if len(folder) > 0 {
		return c.Status(200).JSON(folder)
	}
	return c.JSON(fiber.Map{
		"status":  "False",
		"message": "Failed",
		"error":   "Parse data failed",
		"data":    "null",
	})
}

func UpdateFolder(c *fiber.Ctx) error {
	var folder models.Folder
	var fol []Folder
	if err := c.BodyParser(&folder); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	database.Database.Db.Where("id = ?", folder.Id).Find(&fol)
	for i := 0; i < len(fol); i++ {
		if folder.Id == fol[i].Id {
			database.Database.Db.Debug().Updates(folder)
			db := database.Database.Db.Preload("Request", func(db *gorm.DB) *gorm.DB {
				return database.Database.Db.Preload("Response")
			}).First(&folder)
			if db.Error != nil {
				logrus.Debug("--------------->", db.Error)
			}
			return c.Status(200).JSON(folder)
		}
	}
	return c.JSON(fiber.Map{"status": false})
}
func GetFolderByIdCollection(c *fiber.Ctx) error {
	var collection models.Collection
	var folder []Folder
	var query Query
	c.QueryParser(&query)
	database.Database.Db.Where("id = ?", query.ID).Find(&collection)
	database.Database.Db.Where("collection_id = ?", collection.Id).Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&folder)
	return c.JSON(folder)
}
