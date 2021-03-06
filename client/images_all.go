package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"strconv"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"

	"github.com/parnurzeal/gorequest"
)

// GetAllImages - retrieves all registered Images with params and pagination
// Accepts the CSP struct, page number, pagesize and params map
// Returns Images struct, remaining count and next page
// v2/images
func (cli *Client) GetAllImages(page, pagesize int, paramsString map[string]string, paramsBool map[string]bool) (images.Images, int, int, int) {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response images.Images
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/images"
	paramString := cli.GetStringParams(paramsString)
	paramBool := cli.GetBoolParams(paramsBool)
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).
		Query(paramString).Query(paramBool).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetAllImages from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return response, remaining, next, response.Count
}
