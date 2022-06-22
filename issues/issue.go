package issues

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

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

func sortCommentsById(c []IssueComment) {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Id > c[j].Id
	})
}

func formatDate(date string) string {
	t, err := time.Parse("2006-01-02", strings.Split(date, "T")[0])
	if err != nil {
		log.Fatalf("something went wrong when parsing date: %v", err)
	}

	return t.Format("02.01.06")
}

func printComments(ic IssueComments) {
	sortCommentsById(ic.Comments)
	for _, c := range ic.Comments {
		fmt.Printf("%s by %s\n", formatDate(c.Created), c.Author.DisplayName)
		fmt.Println(c.Body)
		fmt.Println("-------")
	}
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
	printComments(issue.Fields.Comment)
}

func GetIssue(issueKey string) {
	url := helpers.GetIssueUrl(issueKey)
	printIssue(requestIssue(url))
}
