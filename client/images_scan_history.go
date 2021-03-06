package client // import "github.com/BryanKMorrow/aqua-sdk-go/client"

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/BryanKMorrow/aqua-sdk-go/types/images"
	"github.com/parnurzeal/gorequest"
)

// GetScanHistory -  retrieves a single image scan history based on registry, repo and tag
// Path parameters of {registry}, {repo} and {tag}
// Returns response struct
// v2/images/{registry}/{repo}/{tag}/scan_history
func (cli *Client) GetScanHistory(registry, repo, tag string) (images.ScanHistories, error) {
	var response images.ScanHistories
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	request.Set("Authorization", "Bearer "+cli.token)
	apiPath := fmt.Sprintf("/api/v2/images/%s/%s/%s/scan_history", registry, repo, tag)
	events, body, errs := request.Clone().Get(cli.url + apiPath).End()
	if errs != nil {
		log.Println(events.StatusCode)
	}
	if events.StatusCode == 200 {
		err := json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetScanHistory from %s%s, %v ", cli.url, apiPath, err)
		}
	} else {
		return response, errors.New("image not found")
	}
	return response, nil
}
