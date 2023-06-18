package routers

import (
	"core/database"
	"core/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Import(c *fiber.Ctx) error {
	type Info struct {
		Id        *int   `json:"id,omitempty"`
		PostmanId string `json:"_postman_id,omitempty"`
		Name      string `json:"name,omitempty"`
		Schema    string `json:"schema,omitempty"`
		EpID      string `json:"_exporter_id,omitempty"`
	}
	type Url struct {
		Id         *int   `json:"id,omitempty"`
		Raw        string `json:"raw,omitempty"`
		Protocol   string `json:"protocol,omitempty"`
		Host       []byte `json:"host,omitempty"`
		Port       string `json:"port,omitempty"`
		Path       []byte `json:"path,omitempty"`
		ReposesId  *int   `json:"reposes_id,omitempty"`
		RequestId  *int   `json:"request_id,omitempty"`
		Request1Id *int   `json:"request1_id,omitempty"`
	}
	type Raw1 struct {
		Id       int    `json:"id"`
		Language string `json:"language,omitempty"`
		OptionId int    `json:"option_id"`
	}
	type Option struct {
		Id     int    `json:"id"`
		Raw    []Raw1 `json:"raw,omitempty"`
		BodyId int    `json:"body_id"`
	}
	type Reposess struct {
		Id              *int   `json:"id,omitempty"`
		Name            string `json:"name,omitempty"`
		OriginalRequest string `json:"original_request,omitempty"`
		Ppl             string `json:"ppl,omitempty"`
		Hearder         string `json:"hearder,omitempty"`
		Url             []Url  `json:"url,omitempty"`
		RequestId       *int   `json:"request_id,omitempty"`
	}
	type Request struct {
		Id         *int           `json:"id,omitempty"`
		Name       string         `json:"name,omitempty"`
		Method     string         `json:"method,omitempty"`
		Header     string         `json:"header,omitempty"`
		Url        string         `json:"url,omitempty"`
		Cookie     string         `json:"cookie,omitempty"`
		Data       datatypes.JSON `json:"body,omitempty"`
		DataId     *int           `json:"data_id,omitempty"`
		RawId      *int           `json:"raw_id,omitempty"`
		Request1Id *int           `json:"request1_id,omitempty"`
	}
	// type Data struct {
	// 	Id      int       `json:"id"`
	// 	Request []Request `json:"request,omitempty"`
	// }
	// type Body struct {
	// 	Id         int            `json:"id"`
	// 	Mode       string         `json:"mode,omitempty"`
	// 	Raw        datatypes.JSON `json:"raw,omitempty"`
	// 	Option     []Option       `json:"options,omitempty"`
	// 	Request1Id *int           `json:"request1_id"`
	// 	ResponseId *int           `json:"response_id"`
	// }
	type OriginalRequest struct {
		Method     string         `json:"method,omitempty"`
		Hearder    string         `json:"hearder,omitempty"`
		Data       datatypes.JSON `json:"body,omitempty"`
		Url        string         `json:"url,omitempty"`
		ResponseId *int           `json:"response_id,omitempty"`
	}
	type Response struct {
		Id                     *int              `json:"id,omitempty"`
		Name                   string            `json:"name,omitempty"`
		Method                 string            `json:"method,omitempty"`
		OriginalRequest        []OriginalRequest `json:"originalRequest,omitempty"`
		PostmanPreviewLanguage string            `json:"_postman_previewlanguage,omitempty"`
		Header                 string            `json:"header,omitempty"`
		Cookie                 []byte            `json:"cookie,omitempty"`
		Body                   datatypes.JSON    `json:"body,omitempty"`
		Request1Id             *int              `json:"request1_id,omitempty"`
		Url                    string            `json:"url,omitempty"`
		Item_itemId            *int              `json:"item_id,omitempty"`
	}
	type Request1 struct {
		Id          *int           `json:"id,omitempty,-"`
		Name        string         `json:"name,omitempty"`
		Method      string         `json:"method,omitempty"`
		Header      datatypes.JSON `json:"header,omitempty"`
		Body        datatypes.JSON `json:"body,omitempty"`
		Url         string         `json:"url,omitempty"`
		Request     []Request      `json:"request,omitempty"`
		Response    []Response     `json:"response,omitempty"`
		Item_itemId *int           `json:"item_id,omitempty"`
	}
	type Item_item struct {
		Id   *int       `json:"id,omitempty"`
		Name string     `json:"name,omitempty"`
		Item []Request1 `json:"item,omitempty"`
	}
	// type Item struct {
	// 	Id        *int        `json:"id,omitempty"`
	// 	Item_item []Item_item `json:"item,omitempty"`
	// }
	type Ob struct {
		Id   *int        `json:"id,omitempty"`
		InFo Info        `json:"info,"`
		Item []Item_item `json:"item,"`
	}
	type Import struct {
		Data Ob `json:"data"`
	}
	type InFo1 struct {
		PostmanId  string `json:"_postman_id"`
		Name       string `json:"name"`
		Schema     string `json:"schema"`
		ExporterID *int   `json:"_exporter_id,omitempty" gorm:"primaryKey"`
	}
	type Variable1 struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var collection models.Collection
	var folder models.Folder
	var request models.Request
	var response models.Response
	var input Ob
	c.BodyParser(&input)
	Item := input.Item
	InFo := input.InFo
	collection.Name = InFo.Name
	database.Database.Db.Debug().Create(&collection)

	fmt.Println("----len-----", len(Item))
	for i := 0; i < len(Item); i++ {
		if Item[i].Item != nil {
			folder.Name = Item[i].Name
			folder.CollectionId = collection.Id
			database.Database.Db.Debug().Create(&folder)

			for j := 0; j < len(Item[i].Item); j++ {
				// fmt.Println("-----len item-item ----", len(Item[i].Item))
				// fmt.Println("----request----", Item[i].Item)              // request
				// fmt.Println("----response----", Item[i].Item[j].Response) //response

				if Item[i].Item[j].Request != nil {
					request.Name = Item[i].Item[j].Request[j].Name
					//fmt.Println("----name----", Item[i].Item[j].Name)
					request.Method = Item[i].Item[j].Method
					request.Url = Item[i].Item[j].Url
					request.Data = Item[i].Item[j].Body
					request.CollectionId = collection.Id
					request.FolderId = &folder.CollectionId
					database.Database.Db.Debug().Create(&request)

					if Item[i].Item[j].Response != nil {
						for k := 0; k < len(Item[i].Item[j].Response); k++ {
							response.Name = Item[i].Item[j].Response[k].Name
							response.Method = Item[i].Item[j].Response[k].OriginalRequest[k].Method
							response.Url = Item[i].Item[j].Response[k].OriginalRequest[k].Url
							response.Data = Item[i].Item[j].Response[k].OriginalRequest[k].Data
							response.RequestId = request.Id
							response.CollectionId = collection.Id
							database.Database.Db.Debug().Create(&response)
						}
					}
				}
			}
		}
	}
	database.Database.Db.Preload("Folders", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Requests", func(db *gorm.DB) *gorm.DB {
			return database.Database.Db.Preload("Responses")
		})
	}).Preload("Request", func(db *gorm.DB) *gorm.DB {
		return database.Database.Db.Preload("Responses")
	}).Find(&collection)
	return c.JSON(collection)
}
