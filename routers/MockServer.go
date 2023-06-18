package routers

import (
	"core/database"
	"core/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

type Collections struct {
	Id        int         `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name"`
	Folders   []Folders   `json:"folders,omitempty"`
	Request   []Request   `json:"requests,omitempty"`
	Responses []Responses `json:"repsonses,omitempty"`
}
type Folders struct {
	Id           int        `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	Requests     []Requests `json:"requests,omitempty"`
	CollectionId int        `json:"collection_id"`
}
type Requests struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Data         datatypes.JSON `json:"data"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	CollectionId int            `json:"collection_id"`
	FolderId     *int           `json:"folder_id"`
	Responses    []Responses    `json:"repsonses,omitempty"`
}
type Responses struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Method       string         `json:"method"`
	Url          string         `json:"url"`
	Data         datatypes.JSON `json:"data"`
	RequestId    int            `json:"request_id"`
	CollectionId int            `json:"collection_id"`
}
type InFo struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
type ER struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    string `json:"data"`
}
type QueryParams struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var method string

func Method(c *fiber.Ctx) error {
	method := c.Method()
	if method == "GET" {
		return MockServer("GET")(c)
	}
	if method == "POST" {
		return MockServer("POST")(c)
	}
	if method == "PUT" {
		return MockServer("PUT")(c)
	}
	if method == "DELETE" {
		return MockServer("DELETE")(c)
	}
	if method == "PATCH" {
		return MockServer("PATCH")(c)
	}
	return nil
}

func MockServer(method string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		path := c.Params("path")
		var collection models.Collection
		var responses models.Response
		type BodyData struct {
			Data      datatypes.JSON
			RequestId string
		}
		type BodyName struct {
			Name datatypes.JSON
		}
		type Paramsa struct {
			Id        int
			Name      string
			RequestID int
		}
		type DataName struct {
			DataName string
		}

		var input BodyData
		var input_name BodyName
		var input_id Paramsa
		var input_name_param Paramsa
		var input_data_name DataName
		c.BodyParser(&input)
		c.BodyParser(&input_name)
		c.BodyParser(&input_data_name)
		c.ParamsParser(&input_id.Id)
		c.ParamsParser(&input_name_param)
		// lấy id từ params
		paramsa_id := c.FormValue("id")
		paramsa_id = strings.ReplaceAll(paramsa_id, "\"", "")
		// lấy name từ params
		paramsa_name := c.FormValue("name")
		paramsa_name = strings.ReplaceAll(paramsa_name, "\"", "")
		//lấy data từ body
		bodyString := input.Data.String()
		bodyString = strings.ReplaceAll(bodyString, "\"", "")
		// lấy name từ body
		bodyName := input_name.Name.String()
		bodyName = strings.ReplaceAll(bodyName, "\"", "")

		// request_id := input.RequestId
		// request_id = strings.ReplaceAll(request_id, "\"", "")

		request_id_params := c.FormValue("request_id")
		request_id_params = strings.ReplaceAll(request_id_params, "\"", "")

		bodyDataName := input_data_name.DataName
		bodyDataName = strings.ReplaceAll(bodyDataName, "\"", "")
		url := "http://127.0.0.1:8000/mock_server/" + path
		database.Database.Db.Where("name = ?", name).First(&collection)
		//fmt.Printf(request_id)
		//trả về theo id từ params
		if paramsa_id != "" {
			database.Database.Db.Where("method = ?", method).Where("collection_id = ?", collection.Id).Where("url = ?", url).Where("id = ?", paramsa_id).Find(&responses)
			if responses.Name == "" {
				goto res
			}
			return c.JSON(responses.Data)
		}
		//trả về theo name từ paramas
		if paramsa_name != "" {
			database.Database.Db.Where("method = ?", method).Where("collection_id = ?", collection.Id).Where("url = ?", url).Where("name LIKE  ?", "%"+paramsa_name+"%").Find(&responses)
			if responses.Name == "" {
				goto res
			}
			return c.JSON(responses.Data)
		}
		// trả về theo request_id từ body
		if request_id_params != "" {
			database.Database.Db.Where("method = ?", method).Where("collection_id = ?", collection.Id).Where("url = ?", url).Where("request_id = ?", request_id_params).Find(&responses)
			if responses.Name == "" {
				goto res
			}
			return c.JSON(responses.Data)
		}
		//trả về theo name từ body
		if bodyName != "" {
			database.Database.Db.Where("method = ?", method).Where("collection_id = ?", collection.Id).Where("url = ?", url).Where("name LIKE ?", "%"+bodyName+"%").Find(&responses)
			if responses.Data.String() == "" {
				goto res
			}
			return c.JSON(responses.Data)
		}

		if paramsa_name == "" && paramsa_id == "" && bodyName == "" && bodyString == "" && request_id_params == "" {
			database.Database.Db.Where("method = ?", method).Where("collection_id = ?", collection.Id).Where("url = ?", url).Find(&responses)
			return c.JSON(responses.Data)
		}
	res:
		er := []ER{
			{
				Status:  "False",
				Message: "Failed",
				Error:   "Parse data failed",
				Data:    "null"},
		}
		return c.JSON(er)
	}
}
