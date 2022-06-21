package transitions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func getStartIssueUrl(issueKey string) string {
	return fmt.Sprintf("https://%s/rest/api/2/issue/%s/transitions", main.GetJiraSpaceUri(), issueKey)
}

func setIssueToOpen(issueKey string) {
	url := getStartIssueUrl(issueKey)
	req, err := http.NewRequest("POST", url, getTransitionBody(viper.GetInt("transitions.open")))
	if err != nil {
		log.Fatalf("Something went wrong when creating request: %v", err)
	}
	req.Header.Set("Authorization", createBasicToken())
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Something went wrong when trying request: %v", err)
	}

	if res.StatusCode != http.StatusNoContent {
		log.Fatalf("Status is already set")
	} else {
		fmt.Printf("Status %s is set to OPEN", issueKey)
	}

}
