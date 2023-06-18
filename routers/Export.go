package routers

import (
	"core/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func Export2(c *fiber.Ctx) error {
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
		Id                     *int            `json:"id,omitempty"`
		Name                   string          `json:"name,omitempty"`
		Method                 string          `json:"method,omitempty"`
		OriginalRequest        OriginalRequest `json:"originalRequest,omitempty"`
		PostmanPreviewLanguage string          `json:"_postman_previewlanguage,omitempty"`
		Header                 string          `json:"header,omitempty"`
		Cookie                 []byte          `json:"cookie,omitempty"`
		Body                   datatypes.JSON  `json:"body,omitempty"`
		Request1Id             *int            `json:"request1_id,omitempty"`
		Url                    string          `json:"url,omitempty"`
		Item_itemId            *int            `json:"item_id,omitempty"`
	}
	type Request1 struct {
		Id          *int           `json:"id,omitempty,-"`
		Name        string         `json:"name,omitempty"`
		Method      string         `json:"method,omitempty"`
		Header      datatypes.JSON `json:"header,omitempty"`
		Body        datatypes.JSON `json:"body,omitempty"`
		Url         string         `json:"url,omitempty"`
		Request     Request        `json:"request,omitempty"`
		Response    []Response     `json:"response,omitempty"`
		Item_itemId *int           `json:"item_id,omitempty"`
	}
	type Item_item struct {
		Id       *int       `json:"id,omitempty"`
		Name     string     `json:"name,omitempty"`
		Item     []Request1 `json:"item"`
		Requests []Request  `json:"request,omitempty"`
		//Request1 []Request1 `json:"request,omitempty"`
	}
	// type Item struct {
	// 	Id        *int        `json:"id,omitempty"`
	// 	Item_item []Item_item `json:"item,omitempty"`
	// }
	type Ob struct {
		Id   *int        `json:"id,omitempty"`
		InFo Info        `json:"info"`
		Item []Item_item `json:"item"`
		//Item1 []Request1  `json:"item "`
	}
	// type ExportFile struct {
	// 	obj Ob `json:"obj"`
	// }
	type Import struct {
		Data Ob `json:"data"`
	}
	type InFo1 struct {
		PostmanId  string `json:"_postman_id"`
		Name       string `json:"name"`
		Schema     string `json:"schema"`
		ExporterID *int   `json:"_exporter_id,radom,omitempty" gorm:"primaryKey"`
	}
	type Variable1 struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	// type Export1 struct {
	// 	INFO     InFo1     `json:"info"`
	// 	ITEM     Item      `json:"item"`
	// 	VARIABLE Variable1 `json:"variable"`
	// }

	var query Query
	c.QueryParser(&query)
	var collection Collection
	folder := []Folder{}
	request := []Request{}
	response := []Response{}
	items := []Item_item{}
	database.Database.Db.Where("id = ?", query.ID).First(&collection)
	if collection.Id == 0 {
		return c.SendString("sai")
	}
	obj := Ob{
		InFo: Info{
			PostmanId: uuid.New().String(),
			Name:      collection.Name,
			Schema:    "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
			EpID:      query.ExId,
		},
		Item: items,
	}
	rand := rand.Intn(99999999999)
	str := collection.Name + "_" + strconv.Itoa(query.ID) + "_" + strconv.Itoa(rand)
	///home/btdat/DAT/export/DatDemo_2_499379.json
	file, err := os.Create(str + ".json")

	if err != nil {
		return err
	}
	defer file.Close()
	database.Database.Db.Where("collection_id = ?", collection.Id).Find(&folder)
	fmt.Println("----folder-----", len(folder))
	if len(folder) > 0 {
		for i := 0; i < len(folder); i++ {
			fmt.Println("-----folder------", folder)
			new_item_folder := Item_item{
				Name: folder[i].Name,
			}
			database.Database.Db.Where(" folder_id = ?", folder[i].Id).Find(&request)
			if len(request) > 0 {
				for j := 0; j < len(request); j++ {
					new_request_item := Request1{
						Name: request[j].Name,
						Request: Request{
							Method: request[j].Method,
							Header: string(request[j].Header),
							Url:    request[j].Url,
						},
					}
					database.Database.Db.Where("request_id = ?", request[j].Id).Find(&response)
					if len(response) <= 0 {
						new_request_item.Response = append(new_request_item.Response, Response{})
					}
					if len(response) > 0 {
						for k := 0; k < len(response); k++ {
							new_response_item := Response{
								Name:                   response[k].Name,
								PostmanPreviewLanguage: response[k].PostmanPreviewLanguage,
								OriginalRequest: OriginalRequest{
									Method:  response[k].Method,
									Hearder: response[k].Header,
									Url:     response[k].Url,
								},
								Header: response[k].Header,
								Cookie: response[k].Cookie,
								Body:   response[k].Body,
							}
							new_request_item.Response = append(new_request_item.Response, new_response_item)
						}
					}

					new_item_folder.Item = append(new_item_folder.Item, new_request_item)
				}
			}
			if len(request) == 0 {

				new_item_folder.Item = append(new_item_folder.Item, Request1{})
			}
			obj.Item = append(obj.Item, new_item_folder)

			// 	database.Database.Db.Where("collection_id = ?", collection.Id).Find(&request)
			// 	if len(request) > 0 {
			// 		new_item_folder := Item_item{
			// 			Item: []Request1{},
			// 		}
			// 		for j := 0; j < len(request); j++ {
			// 			new_request_item := Request1{
			// 				Name: request[j].Name,
			// 				Request: Request{
			// 					Method: request[j].Method,
			// 					Header: string(request[j].Header),
			// 					Url:    request[j].Url,
			// 				},
			// 			}
			// 			database.Database.Db.Where("request_id = ?", request[j].Id).Find(&response)
			// 			if len(response) <= 0 {
			// 				new_response_item_nil := Response{}
			// 				new_request_item.Response = append(new_request_item.Response, new_response_item_nil)
			// 			}
			// 			if len(response) > 0 {
			// 				for k := 0; k < len(response); k++ {
			// 					new_response_item := Response{
			// 						Name:                   response[k].Name,
			// 						PostmanPreviewLanguage: response[k].PostmanPreviewLanguage,
			// 						OriginalRequest: OriginalRequest{
			// 							Method:  response[k].Method,
			// 							Hearder: response[k].Header,
			// 							Url:     response[k].Url,
			// 						},
			// 						Header: response[k].Header,
			// 						Cookie: response[k].Cookie,
			// 						Body:   response[k].Body,
			// 					}
			// 					new_request_item.Response = append(new_request_item.Response, new_response_item)
			// 				}
			// 			}

			// 			new_item_folder.Item = append(new_item_folder.Item, new_request_item)
			// 		}
			// 		obj.Item = append(obj.Item, new_item_folder)
			// 	}
		}
	}
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	data, _ := json.MarshalIndent(obj, "", "	")
	//data, err := json.Marshal(obj)
	ioutil.WriteFile(str+".json", data, 0644)
	return c.JSON(obj)
}
