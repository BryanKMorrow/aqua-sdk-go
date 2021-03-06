package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"

	"github.com/parnurzeal/gorequest"
)

// GetVulnerabilities - retrieves all vulnerabilities from a particular image
// Accepts the registry, repo and tag strings as well as the page number, pagesize and params map
// Returns response struct, remaining count and next page
// v2/images
func (cli *Client) GetVulnerabilities(registry, repo, tag string, page, pagesize int, paramsString map[string]string, paramsBool map[string]bool) (images.Vulnerabilities, int, int, int) {
	// set the default pagesize
	if pagesize == 0 {
		pagesize = 1000
	}
	var response images.Vulnerabilities
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s/vulnerabilities", registry, repo, tag)
	paramString := cli.GetStringParams(paramsString)
	paramBool := cli.GetBoolParams(paramsBool)
	events, body, errs := request.Clone().Get(cli.url+apiPath).Param("page", strconv.Itoa(page)).Param("pagesize", strconv.Itoa(pagesize)).
		Query(paramString).Query(paramBool).End()
	if errs != nil {
		log.Println(errs)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetVulnerabilities from %s%s, %v ", cli.url, apiPath, err)
			//json: Unmarshal(non-pointer main.Request)
		}
	}
	remaining := cli.CalcRemaining(pagesize, page, response.Count)
	page = response.Page + 1
	remaining, next := cli.CalcNext(remaining, page)
	return response, remaining, next, response.Count
}
