package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/sbvalois/gojira/setup"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type IssueSprintType struct {
}

type IssueSearch struct {
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
	Issues     []struct {
		Id     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Created           string `json:"created"`
			Updated           string `json:"updated"`
			IssueName         string `json:"summary"`
			CustomField_10020 []struct {
				SprintName string `json:"name"`
			} `json:"customfield_10020"`
		} `json:"fields"`
	} `json:"issues"`
}

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Assignee    struct {
			DisplayName string `json:"displayName"`
		} `json:"assignee"`
		Creator struct {
			DisplayName string `json:"displayName"`
		} `json:"creator"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		SubTask []struct {
			Id  string `json:"id"`
			Key string `json:"key"`
		} `json:"subtasks"`
		CustomField_10020 []struct {
			SprintName string `json:"name"`
		} `json:"customfield_10020"`
		Comment struct {
			Comments []struct {
				Id string `json:"id"`
			}
			Total      int `json:"total"`
			MaxResults int `json:"maxResults"`
			StartAt    int `json:"startAt"`
		} `json:"comment"`
	} `json:"fields"`
}

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
	fmt.Println("getJiraSpaceUri", getJiraSpaceUri())
	return fmt.Sprintf(
		"https://%s/rest/api/2/search?jql=status+in+(\"%s\")+and+assignee+=+\"%s\"",
		getJiraSpaceUri(),
		formatIssueType(issueType),
		getJiraEmail(),
	)
}

func createBasicToken() string {
	email := viper.GetString("user.email")
	apiKey := viper.GetString("jira.apiKey")
	encoded := base64.StdEncoding.EncodeToString([]byte(email + ":" + apiKey))

	return "Basic " + encoded
}

// TODO Kan lage ett flag som sorterer på date
func requestIssues(issuesUrl string) IssueSearch {
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\")+and+assignee+=+\"sv@email.24sevenoffice.com\"+order+by+priority"
	req, err := http.NewRequest("GET", issuesUrl, nil)
	if err != nil {
		panic(err)
	}
	//	req.Header.Set("Authorization", "Basic "+b64Key)
	req.Header.Set("Authorization", createBasicToken())
	client := &http.Client{}

	res, err := client.Do(req)
	fmt.Println("res", res)

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

func requestIssue(issueUrl string) Issue {
	req, err := http.NewRequest("GET", issueUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", createBasicToken())
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

func printIssues(issues IssueSearch) {
	//	fmt.Println(PrettyPrint(issues))
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

func printIssue(issue Issue) {
	//	fmt.Println(PrettyPrint(issue))
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

func getJiraSpaceUri() interface{} {
	return viper.Get("jira.space")
}

func getJiraEmail() interface{} {
	return viper.Get("user.email")
}

func getIssues(issueType string) {
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\",+\"to+do\",+\"done\")+and+assignee+=+\"sv@email.24sevenoffice.com\"+order+by+priority"
	url := getIssuesUrl(issueType)
	printIssues(requestIssues(url))

}

func getIssueUrl(issueKey string) string {
	return fmt.Sprintf("https://%s/rest/api/2/issue/%s", getJiraSpaceUri(), issueKey)
}

func getIssue(issueKey string) {
	//	https://24so.atlassian.net/rest/api/2/issue/ERP-2213
	//	var url = "https://24so.atlassian.net/rest/api/2/search?jql=status+in+(\"in+progress\",+\"to+do\",+\"done\")+and+assignee+=+\"
	url := getIssueUrl(issueKey)
	printIssue(requestIssue(url))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getStartIssueUrl(issueKey string) string {
	return fmt.Sprintf("https://%s/rest/api/2/issue/%s/transitions", getJiraSpaceUri(), issueKey)
}

func getTransitionBody(transitionCode int) *bytes.Buffer {
	postBody, _ := json.Marshal(map[string]map[string]int{
		"transition": {
			"id": transitionCode,
		},
	})
	return bytes.NewBuffer(postBody)
}

func setIssueToOpen(issueKey string) {
	url := getStartIssueUrl(issueKey)
	req, err := http.NewRequest("POST", url, getTransitionBody(141))
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

func setIssueToInProgress(issueKey string) {
	url := getStartIssueUrl(issueKey)
	req, err := http.NewRequest("POST", url, getTransitionBody(11))
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
		fmt.Printf("Status %s is set to IN PROGRESS", issueKey)
	}

}

func setIssueToReview(issueKey string) {
	url := getStartIssueUrl(issueKey)
	req, err := http.NewRequest("POST", url, getTransitionBody(31))
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
		fmt.Printf("Issue %s is set to REVIEW", issueKey)
	}
}

func setIssueToQA(issueKey string) {
	url := getStartIssueUrl(issueKey)
	req, err := http.NewRequest("POST", url, getTransitionBody(31))
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
		fmt.Printf("Issue %s is set to REVIEW", issueKey)
	}
}

var filename string = "config.toml"

func init() {
	viper.SetConfigName(filename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("It looks like you've never ran the setup before")
			setup.RunBasicSetup(filename, bufio.NewReader(os.Stdin))
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

}

func main() {
	var language string

	//	os.Exit(0)

	app := &cli.App{
		Name:  "Gojira",
		Usage: "Get your jira tasks, right in your terminal",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Aliases:     []string{"l"},
				Usage:       "language for greeting",
				Destination: &language,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "Setup your jira conf",
				Action: func(c *cli.Context) error {
					setup.RunBasicSetup(filename, bufio.NewReader(os.Stdin))
					return nil
				},
			},
			{
				Name:    "issues",
				Usage:   "get issues",
				Aliases: []string{"i"},
				Subcommands: []*cli.Command{
					{
						Name:  "open",
						Usage: "gets all issues with status todo",
						Action: func(c *cli.Context) error {
							getIssues("open")
							return nil
						},
					},
					{
						Name:  "inprogress",
						Usage: "gets all issues with status in progress",
						Action: func(c *cli.Context) error {
							getIssues("inprogress")
							return nil
						},
					},
				},
			},
			{
				Name:  "issue",
				Usage: "Get one issue",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return errors.New("No issue key provided")
					}
					getIssue(c.Args().First())
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "start",
						Usage: "sets an issue to in progress",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							setIssueToInProgress(c.Args().First())
							return nil
						},
					},
					{
						Name:  "open",
						Usage: "sets an issue to open",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							setIssueToOpen(c.Args().First())
							return nil
						},
					},
					{
						Name:  "review",
						Usage: "sets an issue to review. This will only work if the issue is already in progress",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							setIssueToReview(c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Something went wrong, %v", err)
	}
}
