package routers

import (
	"core/database"
	"core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
)

type Response3 struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	Data         datatypes.JSON `json:"data"`
	RequestId    int            `json:"request_id"`
	CollectionId int            `json:"collection_id"`
}

func GetResponse(c *fiber.Ctx) error {
	response := []models.Response{}
	db := database.Database.Db.Find(&response)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	return c.JSON(response)
}

func CreateResponse(c *fiber.Ctx) error {
	var response models.Response
	var request models.Request
	if err := c.BodyParser(&response); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Debug().Where("request_id = ?", response.RequestId).Find(&request)
	database.Database.Db.Debug().Where("collection_id = ?", request.CollectionId).Create(&response)
	return c.Status(200).JSON(response)
}

func DeleteResponse(c *fiber.Ctx) error {
	var repsonses models.Response
	var query Query
	c.QueryParser(&query)
	database.Database.Db.Where("id = ?", query.ID).Debug().Delete(repsonses)
	return c.Status(200).JSON(fiber.Map{"status": true})
}

func UpdateResponse(c *fiber.Ctx) error {
	var repsonses models.Response
	var res []Response
	if err := c.BodyParser(&repsonses); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Where("id = ?", repsonses.Id).Find(&res)
	for i := 0; i < len(res); i++ {
		if res[i].Id == repsonses.Id {
			database.Database.Db.Debug().Updates(repsonses)
			db := database.Database.Db.Preload("Response").First(&repsonses)
			if db.Error != nil {
				logrus.Debug("--------------->", db.Error)
			}
			return c.Status(200).JSON(repsonses)
		}
	}
	return c.JSON(fiber.Map{
		"status":  "False",
		"message": "Failed",
		"error":   "Parse data failed",
		"data":    "null",
	})
}
func GetResponseById(c *fiber.Ctx) error {
	var query Query
	c.QueryParser(&query)
	var response []models.Response
	db := database.Database.Db.Where("id = ?", query.ID).Find(&response)
	if db.Error != nil {
		logrus.Debug("--------------->", db.Error)
	}
	if len(response) > 0 {

		return c.Status(200).JSON(response)
	}
	return c.JSON(fiber.Map{
		"status":  "False",
		"message": "Failed",
		"error":   "Parse data failed",
		"data":    "null",
	})
}
