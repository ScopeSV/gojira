package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func GetJiraSpaceUri() interface{} {
	return viper.Get("jira.space")
}

func GetJiraEmail() interface{} {
	return viper.Get("user.email")
}

func GetIssueUrl(issueKey string) string {
	return fmt.Sprintf("https://%s/rest/api/2/issue/%s", GetJiraSpaceUri(), issueKey)
}

func GetStartIssueUrl(issueKey string) string {
	return fmt.Sprintf("https://%s/rest/api/2/issue/%s/transitions", GetJiraSpaceUri(), issueKey)
}

func CreateBasicToken() string {
	email := viper.GetString("user.email")
	apiKey := viper.GetString("jira.apiKey")
	encoded := base64.StdEncoding.EncodeToString([]byte(email + ":" + apiKey))

	return "Basic " + encoded
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
