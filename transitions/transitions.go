package transitions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sbvalois/gojira/helpers"
)

func getTransitionBody(transitionId int) *bytes.Buffer {
	b, _ := json.Marshal(map[string]map[string]int{
		"transition": {
			"id": transitionId,
		},
	})
	return bytes.NewBuffer(b)
}

func validateResponse(res *http.Response, issueKey string, trType string) {
	if res.StatusCode != http.StatusNoContent {
		log.Fatalf("Status is already set")
	} else {
		fmt.Printf("Issue %s is set to %s\n", issueKey, trType)
	}
}

func Set(issueKey string, id int, trType string) {
	url := helpers.GetStartIssueUrl(issueKey)

	req, err := http.NewRequest("POST", url, getTransitionBody(id))
	if err != nil {
		log.Fatalf("Something went wrong when creating request: %v", err)
	}
	req.Header.Set("Authorization", helpers.CreateBasicToken())
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Something went wrong when trying request: %v", err)
	}

	validateResponse(res, issueKey, trType)
}
