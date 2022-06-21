package issues

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sbvalois/gojira/helpers"
)

func formatIssueType(issueType string) string {
	switch issueType {
	case "open":
		return "open"
	case "inprogress":
		return "in+progress"
	case "qa":
		return "qa"
	}
	return issueType
}

func getIssuesUrl(issueType string) string {
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\",+\"to+do\",+\"done\")+and+assignee+=+\"
	fmt.Println("GetJiraSpaceUri", helpers.GetJiraSpaceUri())
	return fmt.Sprintf(
		"https://%s/rest/api/2/search?jql=status+in+(\"%s\")+and+assignee+=+\"%s\"",
		helpers.GetJiraSpaceUri(),
		formatIssueType(issueType),
		helpers.GetJiraEmail(),
	)
}

func printIssues(issues IssueSearch) {
	fmt.Println("ISSUES")
	fmt.Println("------")
	fmt.Printf("Total: %v\n", issues.Total)
	fmt.Println("---------------------------------")
	for _, issue := range issues.Issues {
		fmt.Printf("Issue key: %v\n", issue.Key)
		fmt.Printf("Issue: %v\n", issue.Fields.IssueName)
		fmt.Println("---------------------------------")
	}
}

func requestIssues(issuesUrl string) IssueSearch {
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\")+and+assignee+=+\"sv@email.24sevenoffice.com\"+order+by+priority"
	req, err := http.NewRequest("GET", issuesUrl, nil)
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

	var search IssueSearch
	if err := json.Unmarshal(body, &search); err != nil {
		log.Fatalf("Error while unmarshal: %v", err)
	}

	return search
}

func GetIssues(issueType string) {
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\",+\"to+do\",+\"done\")+and+assignee+=+\"sv@email.24sevenoffice.com\"+order+by+priority"
	url := getIssuesUrl(issueType)
	printIssues(requestIssues(url))

}
