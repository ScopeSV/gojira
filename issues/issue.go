package issues

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sbvalois/gojira/helpers"
)

func requestIssue(issueUrl string) Issue {
	req, err := http.NewRequest("GET", issueUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", helpers.CreateBasicToken())
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Something went wrong when trying request: %v", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error while reading body: %v", err)
	}

	var issue Issue
	if err := json.Unmarshal(body, &issue); err != nil {
		log.Fatalf("Error while unmarshal: %v", err)
	}

	return issue

}

func printIssue(issue Issue) {
	fmt.Println("============")
	fmt.Println(issue.Key, "-", issue.Fields.Summary)
	fmt.Println("============")
	fmt.Println("INFO")
	fmt.Println("----")
	fmt.Println("Reporter:", issue.Fields.Creator.DisplayName)
	fmt.Println("Assignee:", issue.Fields.Assignee.DisplayName)
	fmt.Println("Status:", issue.Fields.Status.Name)
	fmt.Println("============")
	fmt.Println("DESCRIPTION")
	fmt.Println("----")
	fmt.Println(issue.Fields.Description)
	fmt.Println("============")
	fmt.Println("COMMENTS")
	fmt.Println("----")
}

func GetIssue(issueKey string) {
	//	https://24so.atlassian.net/rest/api/2/issue/ERP-2213
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\",+\"to+do\",+\"done\")+and+assignee+=+\"
	url := helpers.GetIssueUrl(issueKey)
	printIssue(requestIssue(url))
}
