package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-sdk-go/types/assurance"
	"github.com/parnurzeal/gorequest"
	"log"
)

// GetImageAssurance
func (cli *Client) GetImageAssurance() assurance.Images {
	var response = assurance.Images{}
	request := gorequest.New()
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := "/api/v2/image_assurance"
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	log.Printf("Calling %s%s", cli.url, apiPath)
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetImageAssurance from %s%s, %v ", cli.url, apiPath, err.Error())
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	return response
}
