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

	fmt.Println(isBasicSetupComplete())

	// fmt.Println("To be able to change issue status, you need to enter status transition keys")
	// fmt.Println("The ")
	// , _ := reader.ReadString('\n')

	os.Exit(0)
}
