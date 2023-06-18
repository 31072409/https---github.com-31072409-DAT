package routers

import (
	"core/database"
	"core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Request2 struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Data         datatypes.JSON `json:"data"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	FolderId     *int           `json:"folder_id,omitempty"`
	CollectionId int            `json:"collection_id"`
	Responses    []Response     `json:"repsonses,omitempty"`
}
type Response2 struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	Data         datatypes.JSON `json:"data"`
	RequestId    int            `json:"request_id"`
	CollectionId int            `json:"collection_id"`
}

func GetRequest(c *fiber.Ctx) error {
	requests := []models.Request{}
	db := database.Database.Db.Debug().Preload("Responses").Find(&requests)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	return c.JSON(requests)
}

func CreateRequest(c *fiber.Ctx) error {
	var request models.Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Debug().Create(&request)
	return c.Status(200).JSON(request)
}

func DeleteRequest(c *fiber.Ctx) error {
	var request []Request
	var response []Response
	var query Query
	c.QueryParser(&query)
	database.Database.Db.Where("request_id = ?", query.ID).Find(&response)
	database.Database.Db.Where("id = ?", query.ID).Find(&request)

	if len(request) >= 1 {
		for i := 0; i < len(response); i++ {
			database.Database.Db.Where("request_id = ?", response[i].RequestId).Debug().Delete(response)
		}
		database.Database.Db.Where("id = ?", query.ID).Debug().Delete(request)

		return c.Status(200).JSON(fiber.Map{"status": true})
	}
	return c.Status(200).JSON(fiber.Map{"status": false})

}
func GetRequestByName(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var request []models.Request
	database.Database.Db.Where("name = ? ", query.Name).Preload("Requests", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Response")
	}).Find(&request)
	return c.Status(200).JSON(request)
}
func GetRequestById(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var request []models.Request
	db := database.Database.Db.Where("id = ?", query.ID).Preload("Responses").Find(&request)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	if len(request) > 0 {
		return c.Status(200).JSON(request)
	}
	return c.JSON(fiber.Map{
		"status":  "False",
		"message": "Failed",
		"error":   "Parse data failed",
		"data":    "null",
	})
}
func UpdateRequest(c *fiber.Ctx) error {
	var request models.Request
	var req []Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Where("id = ?", request.Id).Find(&req)
	for i := 0; i < len(req); i++ {
		if req[i].Id == request.Id {
			database.Database.Db.Debug().Updates(request)
			db := database.Database.Db.Preload("Response").First(&request)
			if db.Error != nil {
				logrus.Debug("--------------->", db.Error)
			}
			return c.Status(200).JSON(request)
		}
	}
	return c.JSON(fiber.Map{"status": false})
}
