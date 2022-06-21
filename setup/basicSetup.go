package setup

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func readInput(r InputReader) string {
	// var reader = bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// input := r.
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

type InputReader interface {
	ReadString(delim byte) (string, error)
}

func isBasicSetupComplete() bool {
	return viper.GetString("jira.space") != "" &&
		viper.GetString("user.email") != "" &&
		viper.GetString("jira.apiKey") != ""
}

func RunTransitionSetup(filename string, r InputReader) {
	fmt.Println("In order to change issue status, you need to enter status transition ids.")
	fmt.Println("Unfortunately, there are no super easy way to access your transitions ids.")
	fmt.Println("The easiest way is to do a cURL request on a jira issue which is \"in progress\",")
	fmt.Println("as these issues usually have the most transition states.")
	fmt.Println("")
	fmt.Println("curl --request GET \\")
	fmt.Printf("  --url \"https://%s/rest/api/2/issue/{issue-key}/transitions\" \\\n", viper.GetString("jira.space"))
	fmt.Printf("  --user \"%s:%s\" \\\n", viper.GetString("user.email"), viper.GetString("jira.apiKey"))
	fmt.Println("  --header \"Accept: application/json\" \\")
	fmt.Println("")
	fmt.Println("Gojira supports 3 transition states at the moment: \"open\", \"in progress\", \"review\" and \"done\"")
	fmt.Println("Please enter the transition id for each of these states:")
	fmt.Println("Open: ")
	open := readInput(r)
	if open != "" {
		viper.Set("transitions.open", open)
	}
	fmt.Println("In Progress: ")
	inProgress := readInput(r)
	if inProgress != "" {
		viper.Set("transitions.InProgress", inProgress)
	}
	fmt.Println("Review: ")
	review := readInput(r)
	if review != "" {
		viper.Set("transitions.review", review)
	}
	fmt.Println("Done: ")
	done := readInput(r)
	if done != "" {
		viper.Set("transitions.done", done)
	}

	viper.WriteConfigAs(filename)
	fmt.Println("Transition Setup complete")
	fmt.Println("You may at any point run `gojira setup` to change your settings")
}

func RunBasicSetup(filename string, r InputReader) {
	fmt.Println("Setup")
	fmt.Println("=====")
	fmt.Println("We need to setup your Jira integration")
	fmt.Println("(Leave blank if you don't want to overwrite existing values)")
	fmt.Println("--------------")
	fmt.Println("Enter jira space (ex: foobar.atlassian.net)")

	space := readInput(r)

	if space != "" {
		viper.Set("jira.space", strings.TrimSpace(space))
	}

	fmt.Println("Enter email")

	email := readInput(r)
	if email != "" {
		viper.Set("user.email", strings.TrimSpace(email))
	}

	fmt.Println("Enter API Key (You can generate it in your Jira account settings)")
	apiKey := readInput(r)
	if apiKey != "" {
		viper.Set("jira.apiKey", strings.TrimSpace(apiKey))
	}

	viper.WriteConfigAs(filename)

	if isBasicSetupComplete() {
		fmt.Println("Basic Setup complete")
		fmt.Println("Do you want to setup your Jira transitions now? (y/n)")
		if readInput(r) == "y" {
			RunTransitionSetup(filename, r)
		}
	}

	// fmt.Println("To be able to change issue status, you need to enter status transition keys")
	// fmt.Println("The ")
	// , _ := reader.ReadString('\n')

	os.Exit(0)
}
